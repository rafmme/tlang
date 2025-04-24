package main

import (
	"syscall/js"

	"github.com/rafmme/tlang/editor"
	"github.com/rafmme/tlang/object"
)

func main() {
	parseTlang := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			return "No arguments passed in!"
		}

		code := args[0].String()

		return editor.ParseEditorInput(code)
	})

	js.Global().Set("parseTlang", parseTlang)

	env := object.NewEnvironment()
	executeTlang := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			return "No arguments passed in!"
		}

		code := args[0].String()

		return editor.EvaluateEditorInput(code, env)
	})

	js.Global().Set("executeTlang", executeTlang)

	select {}

}
