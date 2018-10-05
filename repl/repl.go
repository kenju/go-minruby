package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kenju/go-minruby/evaluator"
	"github.com/kenju/go-minruby/lexer"
	"github.com/kenju/go-minruby/object"
	"github.com/kenju/go-minruby/parser"
)

const PROMPT = "(=･ω･=)>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)
		parser := parser.New(lexer)

		program := parser.ParseProgram()
		if len(parser.Errors()) != 0 {
			printParserErrors(out, parser.Errors())
			continue
		}

		for _, st := range program.Statements {
			io.WriteString(out, "[DEBUG]"+st.String()+"\n")
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
