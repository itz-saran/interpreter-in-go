package main

import (
	"fmt"
	"os"

	"example.com/repl"
)

func main() {
	fmt.Printf("Hello! this is monkey programming language!\n")
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
