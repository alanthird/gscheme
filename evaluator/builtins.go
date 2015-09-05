package evaluator

import (
	"fmt"
	e "github.com/alanthird/gscheme/environment"
	t "github.com/alanthird/gscheme/types"
)

func BuildEnvironment() (env *e.Environment) {
	var defFunc func(string, func(interface{}, t.Type) (t.Type, error))
	defFunc = func(name string, fn func(interface{}, t.Type) (t.Type, error)) {
		e.Define(env, &t.Symbol{name}, &t.Builtin{fn})
	}

	env = e.New(nil)

	defFunc("begin", begin)

	/* Pairs */
	defFunc("cons", cons)
	defFunc("car", car)
	defFunc("cdr", cdr)

	/* Arithmetic */
	defFunc("+", add)
	defFunc("-", sub)
	defFunc("=", equal_Number)

	/* General */
	defFunc("eq?", eq)
	defFunc("eqv?", eqv)

	return
}

func begin(env interface{}, args t.Type) (t.Type, error) {
	var (
		result, expr t.Type
		err          error
	)

	for ; args != nil; args, _ = t.Cdr(args) {
		expr, err = t.Car(args)
		if err != nil {
			return nil, err
		}

		//fmt.Printf("%s\n", expr)

		result, err = Eval(env.(*e.Environment), expr)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

////////////////////
// Pair Functions //
////////////////////

func cons(env interface{}, args t.Type) (t.Type, error) {
	a, err := car(nil, args)
	if err != nil {
		return nil, err
	}

	b, err := cdr(nil, args)
	if err != nil {
		return nil, err
	}

	return t.Cons(a, b), nil
}

func car(env interface{}, args t.Type) (t.Type, error) {
	return t.Caar(args)
}

func cdr(env interface{}, args t.Type) (t.Type, error) {
	return t.Cdar(args)
}

////////////////////////////
// arithmetical functions //
////////////////////////////

func add(env interface{}, args t.Type) (t.Type, error) {
	var total int64 = 0

	for ; args != nil; args, _ = t.Cdr(args) {
		arg, _ := t.Car(args)
		n, ok := arg.(*t.Number)
		if !ok {
			return nil, fmt.Errorf("ADD: Not a number: %s", arg)
		}

		total = total + n.Value
	}
	return &t.Number{total}, nil
}

func sub(env interface{}, args t.Type) (t.Type, error) {
	arg, _ := t.Car(args)
	n, ok := arg.(*t.Number)
	if !ok {
		return nil, fmt.Errorf("SUB: Not a number: %s", arg)
	}

	total := n.Value

	for args, _ := t.Cdr(args); args != nil; args, _ = t.Cdr(args) {
		arg, _ := t.Car(args)
		n, ok := arg.(*t.Number)
		if !ok {
			return nil, fmt.Errorf("SUB: Not a number: %s", arg)
		}

		total = total - n.Value
	}
	return &t.Number{total}, nil
}

func equal_Number(env interface{}, args t.Type) (t.Type, error) {
	arg, _ := t.Car(args)
	first, ok := arg.(*t.Number)
	if !ok {
		return nil, fmt.Errorf("=: Not a number: %s", arg)
	}

	for args, _ := t.Cdr(args); args != nil; args, _ = t.Cdr(args) {
		arg, _ := t.Car(args)

		if !t.Eqv(first, arg) {
			return &t.Bool{false}, nil
		}
	}
	return &t.Bool{true}, nil
}

/* General Functions */

func eqv(env interface{}, args t.Type) (t.Type, error) {
	first, _ := t.Car(args)

	for args, _ := t.Cdr(args); args != nil; args, _ = t.Cdr(args) {
		arg, _ := t.Car(args)

		if !t.Eqv(first, arg) {
			return &t.Bool{false}, nil
		}
	}
	return &t.Bool{true}, nil
}

func eq(env interface{}, args t.Type) (t.Type, error) {
	first, _ := t.Car(args)

	for args, _ := t.Cdr(args); args != nil; args, _ = t.Cdr(args) {
		arg, _ := t.Car(args)

		if !t.Eq(first, arg) {
			return &t.Bool{false}, nil
		}
	}
	return &t.Bool{true}, nil
}
