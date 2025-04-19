package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/rafmme/tlang/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Print(repl.TLANG_TEXT_LOGO)
	fmt.Printf("\nHallo %s! This is the TLang Programming Language!\n", user.Username)
	fmt.Print("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
