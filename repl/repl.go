package repl

import (
	"bufio"
	"fmt"
	"github.com/kenju/go-minruby/lexer"
	"io"
)

const PROMPT = "(=･ω･=)>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)

		io.WriteString(out, line)
		io.WriteString(out, "\n")
	}
}
