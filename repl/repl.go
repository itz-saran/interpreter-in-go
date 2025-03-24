package repl

import (
	"bufio"
	"fmt"
	"io"

	"example.com/evaluator"
	"example.com/lexer"
	"example.com/object"
	"example.com/parser"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

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

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "---ERROR(S) ENCOUNTERED---\n")
	for i, e := range errors {
		io.WriteString(out, fmt.Sprintf("%d.", i+1)+"\t"+e+"\n")
	}
}
