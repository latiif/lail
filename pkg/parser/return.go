package parser

import (
	"github.com/latiif/lail/pkg/ast"
	"github.com/latiif/lail/pkg/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currToken}

	p.nextToken()
	stmt.ReturnValue = p.parseExpression(Lowest)

	for p.peekTokenIs(token.Semicolon) {
		p.nextToken()
	}
	return stmt
}
