package evaluator

import (
	"fmt"
	e "github.com/alanthird/gscheme/environment"
	"github.com/alanthird/gscheme/types"
)

func buildEnvironment() (env *e.Environment) {
	env = e.New(nil)

	e.Define(env, &types.Symbol{"begin"}, &types.Builtin{begin})

	/* Arithmetic */
	e.Define(env, &types.Symbol{"+"}, &types.Builtin{add})
	e.Define(env, &types.Symbol{"-"}, &types.Builtin{sub})
	e.Define(env, &types.Symbol{"="}, &types.Builtin{equal_Number})

	return
}

func begin(env interface{}, args types.SchemeType) (types.SchemeType, error) {
	var (
		result, expr types.SchemeType
		err          error
	)

	for ; args != nil; args, _ = types.Cdr(args) {
		expr, err = types.Car(args)
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

/* arithmetical functions */

func add(env interface{}, args types.SchemeType) (types.SchemeType, error) {
	var total int64 = 0

	for ; args != nil; args, _ = types.Cdr(args) {
		arg, _ := types.Car(args)
		n, ok := arg.(*types.Number)
		if !ok {
			return nil, fmt.Errorf("ADD: Not a number: %s", arg)
		}

		total = total + n.Value
	}
	return &types.Number{total}, nil
}

func sub(env interface{}, args types.SchemeType) (types.SchemeType, error) {
	arg, _ := types.Car(args)
	n, ok := arg.(*types.Number)
	if !ok {
		return nil, fmt.Errorf("SUB: Not a number: %s", arg)
	}

	total := n.Value

	for args, _ := types.Cdr(args); args != nil; args, _ = types.Cdr(args) {
		arg, _ := types.Car(args)
		n, ok := arg.(*types.Number)
		if !ok {
			return nil, fmt.Errorf("SUB: Not a number: %s", arg)
		}

		total = total - n.Value
	}
	return &types.Number{total}, nil
}

func equal_Number(env interface{}, args types.SchemeType) (types.SchemeType, error) {
	arg, _ := types.Car(args)
	first, ok := arg.(*types.Number)
	if !ok {
		return nil, fmt.Errorf("=: Not a number: %s", arg)
	}

	for args, _ := types.Cdr(args); args != nil; args, _ = types.Cdr(args) {
		arg, _ := types.Car(args)

		if !types.Eqv(first, arg) {
			return &types.Bool{false}, nil
		}
	}
	return &types.Bool{true}, nil
}
