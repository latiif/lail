package interpretor

import (
	"fmt"

	"github.com/latiif/lail/pkg/ast"

	"github.com/latiif/lail/pkg/object"
)

var (
	// True is the constant true
	True = &object.Boolean{Value: true}
	// False is the constant false
	False = &object.Boolean{Value: false}
	// Null is the constant null
	Null = &object.Null{}
)

// Eval recursively evaluates a node
func Eval(node ast.Node, env *object.Env) object.Object {
	switch node := node.(type) {
	case *ast.ImportStatement:
		return evalProgram(node.Program, env)
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.LetStatement:
		rhs := Eval(node.Value, env)
		return env.Set(node.Name.Value, rhs)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.ExpressionStatement:
		res := Eval(node.Expression, env)
		if res.Type() == object.ErrorObject {
			fmt.Println("Compile-time error")
			return nil
		}
		return res
	case *ast.IntegerLiteral:
		return &object.Integer{
			Value: node.Value,
		}
	case *ast.StringLiteral:
		return &object.String{
			Value: node.Value,
		}
	case *ast.Boolean:
		return getBooleanObject(node.Value)
	case *ast.IfExpression:
		if evalAsBoolean(Eval(node.Condition, env)) {
			return Eval(node.Consequence, env)
		}
		if node.Alternative != nil {
			return Eval(node.Alternative, env)
		}
		return Null
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		return &object.Return{
			Value: val,
		}
	case *ast.PrefixExpression:
		rhs := Eval(node.Right, env)
		return evalPrefixExpression(node.Operator, rhs)
	case *ast.InfixExpression:
		rhs := Eval(node.Right, env)
		lhs := Eval(node.Left, env)
		return evalInfixExpression(lhs, node.Operator, rhs)
	case *ast.FunctionLiteral:
		params := node.Params
		body := node.Body
		return &object.Function{
			Params: params,
			Body:   body,
			Env:    env,
		}
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		args := evalExpressions(node.Args, env)
		return applyFunction(function, args)
	}
	return nil
}

// value of a block of statements, is its latest expression value
func evalProgram(prog *ast.Program, e *object.Env) object.Object {
	var result object.Object

	for _, stmt := range prog.Statements {
		result = Eval(stmt, e)
		if result == nil {
			return &object.Error{
				Message: "NULL encounterd.",
			}
		}
		if result.Type() == object.ReturnObject {
			return result.(*object.Return).Value
		}
	}

	return result
}

func getBooleanObject(val bool) *object.Boolean {
	if val {
		return True
	}
	return False
}

func evalPrefixExpression(operator string, operand object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperator(operand)
	case "-":
		return evalMinusOperator(operand)
	default:
		return Null
	}
}

func evalBangOperator(operand object.Object) object.Object {
	if operand.Type() == object.IntegerObject {
		operand = getBooleanObject(operand.(*object.Integer).Value != 0)
	}
	switch operand {
	case True:
		return False
	case False:
		return True
	case Null:
		return True
	default:
		return False
	}
}

func evalMinusOperator(operand object.Object) object.Object {
	if operand.Type() != object.IntegerObject {
		return Null
	}

	val := operand.(*object.Integer).Value
	return &object.Integer{
		Value: -val,
	}
}

func evalInfixExpression(lOperand object.Object, operator string, rOperand object.Object) object.Object {

	if lOperand.Type() != object.IntegerObject && rOperand.Type() == object.IntegerObject && operator == "-" {
		return newIncompatibleTypes(operator, lOperand, rOperand)
	}

	lValue := evalAsInteger(lOperand)
	rValue := evalAsInteger(rOperand)

	switch operator {
	case "+":
		return evalInfixPlus(lOperand, rOperand)
	case "-":
		return evalInfixMinus(lOperand, rOperand)
	case "*":
		return &object.Integer{Value: lValue * rValue}
	case "/":
		// division by zero
		if rValue == 0 {
			return Null
		}
		return &object.Integer{Value: lValue / rValue}
	case ">":
		return getBooleanObject(lValue > rValue)
	case "<":
		return getBooleanObject(lValue < rValue)
	case ">=":
		return getBooleanObject(lValue >= rValue)
	case "<=":
		return getBooleanObject(lValue <= rValue)
	case "!=":
		return evalInfixInequality(lOperand, rOperand)
	case "==":
		return evalInfixEquality(lOperand, rOperand)

	default:
		return Null
	}
}

func evalAsBoolean(operand object.Object) bool {
	switch operand.Type() {
	case object.BooleanObject:
		return operand.(*object.Boolean).Value
	case object.IntegerObject:
		return operand.(*object.Integer).Value != 0
	default:
		return false
	}
}

func evalBlockStatement(block *ast.BlockStatement, e *object.Env) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement, e)
		if result != nil && result.Type() == object.ReturnObject {
			return result
		}
	}
	return result
}

func evalIdentifier(ident *ast.Identifier, e *object.Env) object.Object {
	val, ok := e.Get(ident.Value)
	if !ok {
		return Null
	}
	return val
}
func evalExpressions(exprs []ast.Expression, e *object.Env) []object.Object {
	res := make([]object.Object, len(exprs))

	for i, expr := range exprs {
		res[i] = Eval(expr, e)
	}
	return res
}
func applyFunction(fn object.Object, args []object.Object) object.Object {
	function, ok := fn.(*object.Function)
	if !ok {
		return Null
	}

	fnExtendedEnv := extendFunctionEnv(function, args)
	return unwrapReturnValue(Eval(function.Body, fnExtendedEnv))
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Env {
	env := object.NewEnclosedEnv(fn.Env)
	for i, param := range fn.Params {
		env.Set(param.Value, args[i])
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.Return); ok {
		return returnValue.Value
	}
	return obj
}

func newIncompatibleTypes(operator string, lhs, rhs object.Object) object.Object {
	return &object.Error{
		Message: fmt.Sprintf("Operator %s does not support operands of type %q and %q.", operator, lhs.Type(), rhs.Type()),
	}
}
