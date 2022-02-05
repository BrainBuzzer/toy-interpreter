package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, welcome to BrainBuzzer's toy interpreter.\n", user.Username)
	fmt.Printf("Type in any commands that are supported\n")
	repl.Start(os.Stdin, os.Stdout)
}
