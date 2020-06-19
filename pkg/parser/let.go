package parser

import (
	"github.com/latiif/lail/pkg/ast"
	"github.com/latiif/lail/pkg/token"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currToken}

	if !p.expectPeek(token.Ident) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	if !p.expectPeek(token.Assign) {
		return nil
	}
	p.nextToken()

	stmt.Value = p.parseExpression(Lowest)

	for p.peekTokenIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}
