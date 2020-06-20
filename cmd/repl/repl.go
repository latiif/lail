package repl

import (
	"bufio"
	"bytes"
	"io"

	"github.com/latiif/lail/pkg/evaluator/interpretor"
	"github.com/latiif/lail/pkg/lexer"
	"github.com/latiif/lail/pkg/object"
	"github.com/latiif/lail/pkg/parser"
)

// Prompt is the symbol printed at the beginning of every line
const Prompt = "> "

// Start starts the interactive REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnv()
	for scanner.Scan() {
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l, "./")

		prog := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		interpreted := interpretor.Eval(prog, env)

		if interpreted != nil {
			io.WriteString(out, interpreted.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, err := range errors {
		io.WriteString(out, "\t"+err+"\n")
	}
}

func InterpretFile(context string, in io.Reader, out io.Writer, err io.Writer) {
	scanner := bufio.NewScanner(in)
	var b bytes.Buffer
	for scanner.Scan() {
		b.WriteString(scanner.Text())
		b.WriteString("\n") // to preserve new lines for token logging
	}

	e := object.NewEnv()
	l := lexer.New(b.String())
	p := parser.New(l, context)

	prog := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(err, p.Errors())
		return
	}

	interpreted := interpretor.Eval(prog, e)

	if interpreted != nil {
		io.WriteString(out, interpreted.Inspect())
		io.WriteString(out, "\n")
	}
}
