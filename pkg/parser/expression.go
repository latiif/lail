package parser

import (
	"fmt"
	"strconv"

	"github.com/latiif/lail/pkg/ast"
	"github.com/latiif/lail/pkg/token"
)

const (
	_ int = iota
	Lowest
	Equals
	LessGreater
	Sum
	Product
	Prefix
	Call
)

var precedences = map[token.Type]int{
	token.EQ:       Equals,
	token.NEQ:      Equals,
	token.LT:       LessGreater,
	token.GT:       LessGreater,
	token.GTE:      LessGreater,
	token.LTE:      LessGreater,
	token.Plus:     Sum,
	token.Minus:    Sum,
	token.Slash:    Product,
	token.Astersik: Product,
	token.Lparen:   Call,
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currToken}

	stmt.Expression = p.parseExpression(Lowest)

	if p.peekTokenIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precdence int) ast.Expression {
	prefix := p.prefixParseFns[p.currToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.Semicolon) && precdence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.currToken}

	val, err := strconv.ParseInt(p.currToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = val

	return lit
}

func (p *Parser) parseStringLiteral() ast.Expression {
	lit := &ast.StringLiteral{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}
	return lit
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
	}
	p.nextToken()
	expr.Right = p.parseExpression(Prefix)

	return expr
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
		Left:     left,
	}
	precedence := p.currPrecedence()
	p.nextToken()
	expr.Right = p.parseExpression(precedence)

	return expr
}

func (p *Parser) noPrefixParseFnError(t token.Type) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return Lowest
}

func (p *Parser) currPrecedence() int {
	if p, ok := precedences[p.currToken.Type]; ok {
		return p
	}
	return Lowest
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: p.currToken,
		Value: p.currTokenIs(token.True), // if it is not true, it's false
	}
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(Lowest)

	if !p.expectPeek(token.Rparen) {
		return nil
	}

	return exp
}

func (p *Parser) parseIfExpression() ast.Expression {
	exp := &ast.IfExpression{
		Token: p.currToken,
	}
	if !p.expectPeek(token.Lparen) {
		return nil
	}
	p.nextToken()
	exp.Condition = p.parseExpression(Lowest)
	if !p.expectPeek(token.Rparen) {
		return nil
	}

	exp.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(token.Else) {
		p.nextToken()
		exp.Alternative = p.parseBlockStatement()
	}

	return exp
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	if !p.expectPeek(token.Lbrace) {
		return nil
	}

	block := &ast.BlockStatement{
		Token: p.currToken,
	}

	p.nextToken()

	block.Statements = []ast.Statement{}

	for !p.currTokenIs(token.Rbrace) && !p.currTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	exp := &ast.FunctionLiteral{
		Token: p.currToken,
	}
	if !p.expectPeek(token.Lparen) {
		return nil
	}

	exp.Params = p.parseFunctionParams()

	exp.Body = p.parseBlockStatement()

	return exp
}

func (p *Parser) parseFunctionParams() []*ast.Identifier {
	params := []*ast.Identifier{}

	// empty params
	if p.peekTokenIs(token.Rparen) {
		p.nextToken()
		return params
	}

	p.nextToken()

	param := &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	params = append(params, param)

	for p.peekTokenIs(token.Comma) {
		p.nextToken() // consume the comma
		p.nextToken()
		param := &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
		params = append(params, param)
	}

	if !p.expectPeek(token.Rparen) {
		return nil
	}

	return params
}

func (p *Parser) parseCallExpression(fnLiteral ast.Expression) ast.Expression {
	exp := &ast.CallExpression{
		Token:    p.currToken,
		Function: fnLiteral,
	}
	exp.Args = p.parseFunctionArgs()
	return exp
}

func (p *Parser) parseFunctionArgs() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(token.Rparen) {
		p.nextToken()
		return args
	}
	p.nextToken()
	arg := p.parseExpression(Lowest)
	args = append(args, arg)

	for p.peekTokenIs(token.Comma) {
		p.nextToken() // consume the comma
		p.nextToken()
		arg := p.parseExpression(Lowest)
		args = append(args, arg)
	}

	if !p.expectPeek(token.Rparen) {
		return nil
	}

	return args
}
