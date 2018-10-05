package repl

import (
	"bufio"
	"fmt"
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

		io.WriteString(out, line)
		io.WriteString(out, "\n")
	}
}