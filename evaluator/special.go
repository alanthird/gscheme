package evaluator

import (
	"fmt"
	"github.com/alanthird/gscheme/environment"
	"github.com/alanthird/gscheme/types"
)

func isSpecialForm(o types.SchemeType) bool {
	if f, ok := o.(*types.Symbol); ok {
		switch f.Value {
		case "define", "quote", "lambda":
			return true
		}
	}
	return false
}

func applySpecialForm(env *environment.Environment, name string, args types.SchemeType) (types.SchemeType, error) {
	switch name {
	case "quote":
		return types.Car(args)
	case "define":
		return nil, define(env, args)
	case "lambda":
		return lambda(env, args)
	}
	return nil, nil
}

func define(env *environment.Environment, args types.SchemeType) error {
	name, err := types.Car(args)
	if err != nil {
		return err
	}
	value, err := types.Cadr(args)
	if err != nil {
		return err
	}

	evaldValue, err := Eval(env, value)
	if err != nil {
		return err
	}
	
	environment.Define(env, name.(*types.Symbol), evaldValue)
	return nil
}

func lambda(env *environment.Environment, a types.SchemeType) (types.SchemeType, error) {
	args, err := types.Car(a)
	if err != nil {
		return nil, fmt.Errorf("%s\nLAMBDA: unable to get argument list: %s", err, a)
	}

	text, err := types.Cdr(a)
	if err != nil {
		return nil, fmt.Errorf("%s\nLAMBDA: unable to get function body: %s", err, a)
	}

	return &types.SFunction{text, args.(*types.Pair), env}, nil
}
