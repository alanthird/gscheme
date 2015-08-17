package evaluator

import (
	"github.com/alanthird/gscheme/environment"
	"github.com/alanthird/gscheme/parser"
	"github.com/alanthird/gscheme/types"
	"strings"
	"testing"
)

func TestEval(t *testing.T) {
	env := environment.New(nil)

	environment.Define(env, &types.Symbol{"everything"}, &types.Number{42})
	environment.Define(env, &types.Symbol{"cows"}, &types.String{"moo"})
	environment.Define(env, &types.Symbol{"sheep"}, &types.String{"baa"})
	environment.Define(env, &types.Symbol{"+"}, &types.Builtin{types.Add})

	if r, _ := Eval(env, &types.Symbol{"everything"}); !types.Eqv(r, &types.Number{42}) {
		t.Error("Var 'everything' != 42: ", r)
	}

	f, _ := parser.Parse(strings.NewReader("(+ 1 1)"))
	if r, err := Eval(env, f); err != nil {
		t.Error("Evaluating (+ 1 1) caused an error: ", err)
	} else if !types.Eqv(r, &types.Number{2}) {
		t.Error("(+ 1 1) != 2: ", r)
	}

	f, _ = parser.Parse(strings.NewReader("(+ 1 (+ 2 3))"))
	if r, err := Eval(env, f); err != nil {
		t.Error("Evaluating (+ 1 (+ 2 3)) caused an error: ", err)
	} else if !types.Eqv(r, &types.Number{6}) {
		t.Error("(+ 1 (+ 2 3)) != 6: ", r)
	}
}

func TestBegin(t *testing.T) {
	env := environment.New(nil)

	environment.Define(env, &types.Symbol{"cows"}, &types.String{"moo"})
	environment.Define(env, &types.Symbol{"sheep"}, &types.String{"baa"})
	environment.Define(env, &types.Symbol{"begin"}, &types.Builtin{Begin})

	f, _ := parser.Parse(strings.NewReader("(begin cows sheep)"))
	if r, err := Eval(env, f); err != nil {
		t.Error("Evaluating (begin cows sheep) caused an error: ", err)
	} else if !types.Eqv(r, &types.String{"baa"}) {
		t.Error("(begin cows sheep) != baa", r)
	}
}

func TestDefine(t *testing.T) {
	env := environment.New(nil)

	f, _ := parser.Parse(strings.NewReader("(define answer 42)"))
	if _, err := Eval(env, f); err != nil {
		t.Error("Evaluating (define answer 42) failed:", err)
	}
	
	if a, err := environment.Get(env, &types.Symbol{"answer"}); err != nil {
		t.Error("'Get'ting 'answer' caused an error: ", err)
	} else if !types.Eqv(a, &types.Number{42}) {
		t.Error("answer != 42", a)
	}
}

func TestQuote(t *testing.T) {
	env := environment.New(nil)

	f, _ := parser.Parse(strings.NewReader("(quote moo)"))
	if a, err := Eval(env, f); err != nil {
		t.Error("Evaluating (quote moo) failed:", err)
	} else if !types.Eqv(a, &types.Symbol{"moo"}) {
		t.Error("(quote moo) != moo", a)
	}
}

func TestLambda(t *testing.T) {
	env := environment.New(nil)

	environment.Define(env, &types.Symbol{"+"}, &types.Builtin{types.Add})
	environment.Define(env, &types.Symbol{"begin"}, &types.Builtin{Begin})

	f, _ := parser.Parse(strings.NewReader("(lambda (n) n)"))
	if a, err := Eval(env, f); err != nil {
		t.Error("Evaluating lambda failed:", err, f)
	} else if !types.IsSFunction(a) {
		t.Error("Evaluating lambda didn't return an SFunction:", a)
	}
	
	f, _ = parser.Parse(strings.NewReader("(define add1 (lambda (n) (+ 1 n)))"))
	if _, err := Eval(env, f); err != nil {
		t.Error("Evaluating define lambda failed:", err)
	}

	if a, err := environment.Get(env, &types.Symbol{"add1"}); err != nil {
		t.Error("'Get'ting add1 caused an error: ", err)
	} else if !types.IsSFunction(a) {
		t.Error("add1 is not an SFunction:", a)
	}
	
	f, _ = parser.Parse(strings.NewReader("(add1 1)"))
	if a, err := Eval(env, f); err != nil {
		t.Error("Evaluating (add1 1) failed:", err)
	} else if !types.Eqv(a, &types.Number{2}) {
		t.Error("(add1 1) != 2", a)
	}
}

func TestScoping(t *testing.T) {
	env := environment.New(nil)

	environment.Define(env, &types.Symbol{"begin"}, &types.Builtin{Begin})
	environment.Define(env, &types.Symbol{"+"}, &types.Builtin{types.Add})

	funcString := "(begin (define a 0) (define get-num ((lambda (a b) (define c 4) (lambda (b) (+ a b c))) 1 2)))"
	
	f, err := parser.Parse(strings.NewReader(funcString))
	if err != nil {
		t.Error("Parsing funcString failed:", err)
	}
	
	if _, err := Eval(env, f); err != nil {
		t.Error("Evaluating funcString failed:", err)
	}

	f, _ = parser.Parse(strings.NewReader("a"))
	if a, err := Eval(env, f); err != nil {
		t.Error("Evaluating 'a' failed:", err)
	} else if !types.Eqv(a, &types.Number{0}) {
		t.Error("a != 0", a)
	}

	f, _ = parser.Parse(strings.NewReader("(get-num 8)"))
	if a, err := Eval(env, f); err != nil {
		t.Error("Evaluating (get-num) failed:", err)
	} else if !types.Eqv(a, &types.Number{13}) {
		t.Error("(get-num 8) != 13:", a)
	}
}
