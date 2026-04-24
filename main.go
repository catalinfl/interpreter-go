package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/catalinfl/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s\n", user.Username)
	fmt.Printf("REPL init\n")
	repl.Start(os.Stdin, os.Stdout)
}
