package repl

import (
	"bufio"
	"fmt"
	"io"

	"example.com/lexer"
	"example.com/parser"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			io.WriteString(out, "Exiting...\n")
			return
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			return
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "---ERROR(S) ENCOUNTERED---\n")
	for i, e := range errors {
		io.WriteString(out, fmt.Sprintf("%d.", i+1)+"\t"+e+"\n")
	}
}
