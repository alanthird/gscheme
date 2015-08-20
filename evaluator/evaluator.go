package evaluator

import (
	"fmt"
	"github.com/alanthird/gscheme/environment"
	"github.com/alanthird/gscheme/types"
)

func Eval(env *environment.Environment, f types.SchemeType) (types.SchemeType, error) {
	if types.IsPair(f) {
		car, err := types.Car(f)
		if err != nil {
			return nil, err
		}

		cdr, err := types.Cdr(f)
		if err != nil {
			return nil, err
		}

		cdrAgain, ok := cdr.(*types.Pair)
		if !ok {
			return nil, fmt.Errorf("EVAL: Not pair: %s", cdr)
		}

		return Apply(env, car, cdrAgain)
	}

	if types.IsSymbol(f) {
		return environment.Get(env, f.(*types.Symbol))
	}

	return f, nil
}

func Apply(env *environment.Environment, f types.SchemeType, args types.SchemeType) (types.SchemeType, error) {
	if isSpecialForm(f) {
		return applySpecialForm(env, f.(*types.Symbol).Value, args)
	}

	fn, err := Eval(env, f)
	if err != nil {
		return nil, fmt.Errorf("%s\nAPPLY: eval: %s", err, f)
	}

	args, err = evalArgs(env, args)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("%s\n", args)

	if bfn, ok := fn.(*types.Builtin); ok {
		if r, err := bfn.Fn(env, args); err != nil {
			return nil, fmt.Errorf("%s\nAPPLY: builtin %s", err, f.(*types.Symbol).Value)
		} else {
			return r, nil
		}
	}

	if sfn, ok := fn.(*types.SFunction); ok {
		sfnEnv := environment.New(sfn.Env.(*environment.Environment))
		err = environment.AddArgs(sfnEnv, sfn.Args, args)
		if err != nil {
			return nil, err
		}

		if r, err := begin(sfnEnv, sfn.Function); err != nil {
			return nil, fmt.Errorf("%s\nAPPLY: function %s", err, f.(*types.Symbol).Value)
		} else {
			return r, nil
		}
	}
	return nil, fmt.Errorf("APPLY: Unknown function: %s", f.(*types.Symbol).Value)
}

func evalArgs(env *environment.Environment, args types.SchemeType) (types.SchemeType, error) {
	if args == nil {
		return nil, nil
	}

	arg, err := types.Car(args)
	if err != nil {
		return nil, err
	}

	car, err := Eval(env, arg)
	if err != nil {
		return nil, err
	}

	next, err := types.Cdr(args)
	if err != nil {
		return nil, err
	}

	cdr, err := evalArgs(env, next)
	if err != nil {
		return nil, err
	}

	return types.Cons(car, cdr), nil
}

func listToArray(env *environment.Environment, list *types.Pair) (argList []types.SchemeType, err error) {
	for a := list; a != nil; {
		car, err := types.Car(a)
		if err != nil {
			return nil, err
		}

		r, err := Eval(env, car)
		if err != nil {
			return nil, err
		}
		argList = append(argList, r)

		tempA, err := types.Cdr(a)
		if err != nil {
			return nil, err
		}

		var ok bool
		a, ok = tempA.(*types.Pair)
		if !ok && a != nil {
			return nil, fmt.Errorf("Eval: arguments not list: %s", tempA)
		}
	}

	return
}
