package parser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/latiif/lail/pkg/ast"
	"github.com/latiif/lail/pkg/lexer"
	"github.com/latiif/lail/pkg/token"
)

func (p *Parser) parseImportStatement() *ast.ImportStatement {
	stmt := &ast.ImportStatement{Token: p.currToken}

	if !p.expectPeek(token.String) {
		return nil
	}

	rawFile, err := retrieveFile(filepath.Join(p.Context, p.currToken.Literal))
	if err != nil {
		rawFile, err = downloadFile(p.currToken.Literal)
	}

	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("Unable to locate and read file at %s", p.currToken.Literal))
		return nil
	}
	li := lexer.New(rawFile)
	pi := New(li, filepath.Dir(filepath.Join(p.Context, p.currToken.Literal)))

	stmt.Program = pi.ParseProgram()

	return stmt
}

func retrieveFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func downloadFile(path string) (string, error) {
	resp, err := http.Get(path)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	return string(contents), nil
}
