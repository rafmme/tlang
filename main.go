package main

import (
	"syscall/js"

	"github.com/rafmme/tlang/editor"
)

func main() {
	executeTlang := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			return "No arguments passed in!"
		}

		code := args[0].String()

		return editor.ParseEditorInput(code)
	})

	js.Global().Set("executeTlang", executeTlang)

	select {}

	/* 	user, err := user.Current()

	   	if err != nil {
	   		panic(err)
	   	}

	   	fmt.Print(repl.TLANG_TEXT_LOGO)
	   	fmt.Printf("\nHallo %s! This is the TLang Programming Language!\n", user.Username)
	   	fmt.Print("Feel free to type in commands\n")
	   	repl.Start(os.Stdin, os.Stdout) */

}
