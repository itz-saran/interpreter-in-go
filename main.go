package main

import (
	"fmt"
	"os"
	"os/user"

	"example.com/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! this is monkey programming language!\n", user)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
