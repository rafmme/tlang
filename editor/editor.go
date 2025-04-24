package editor

import (
	"github.com/rafmme/tlang/evaluator"
	"github.com/rafmme/tlang/lexer"
	"github.com/rafmme/tlang/object"
	"github.com/rafmme/tlang/parser"
	"github.com/rafmme/tlang/repl"
)

func workOnCodeEditorParserErrors(errors []string) string {
	errorMsg := repl.TLANG_TEXT_LOGO
	errorMsg += "\nWoops! We ran into some tlang issues here!\n"
	errorMsg += " parser errors:\n"

	for _, msg := range errors {
		errorMsg += "\t" + msg + "\n"
	}

	return errorMsg
}

func ParseEditorInput(code string) string {
	l := lexer.New(code)
	p := parser.New(l)

	p.ParseProgram()

	if len(p.Errors()) != 0 {
		return workOnCodeEditorParserErrors(p.Errors())
	}

	return "No Error!"
}

func EvaluateEditorInput(code string, env *object.Environment) string {
	l := lexer.New(code)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		return workOnCodeEditorParserErrors(p.Errors())
	}

	evaluated := evaluator.Eval(program, env)

	if evaluated != nil {
		return evaluated.Inspect() + "\n"
	}

	return "No output given!"
}
