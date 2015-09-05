package main

import (
	"fmt"
	"github.com/alanthird/gscheme/evaluator"
	"github.com/alanthird/gscheme/parser"
	"github.com/bobappleyard/readline"
	"io"
	"os"
	"strings"
)

func main() {
	env := evaluator.BuildEnvironment()

	for {
		l, err := readline.String("> ")
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "%s\n", err)
			break
		}

		f, err := parser.Parse(strings.NewReader(l))
		if err != nil {
			fmt.Fprintln(os.Stderr, "%s", err)
			readline.AddHistory(l)
		}

		r, err := evaluator.Eval(env, f)
		if err != nil {
			fmt.Fprintln(os.Stderr, "%s", err)
		} else {
			fmt.Println(r)
		}
		readline.AddHistory(l)
	}
}
