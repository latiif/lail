package parser

import (
	"github.com/latiif/lail/pkg/ast"
	"github.com/latiif/lail/pkg/token"
)

func (p *Parser) parseImportStatement() *ast.ImportStatement {
	stmt := &ast.ImportStatement{Token: p.currToken}

	if !p.expectPeek(token.String) {
		return nil
	}

	return stmt
}
