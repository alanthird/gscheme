package environment

import (
	"fmt"
	"github.com/alanthird/gscheme/types"
)

type envError struct {
	m string
	e error
}

func (e *envError) Error() string {
	if e.e != nil {
		return fmt.Sprintf("%s\nENV: %s", e.e, e.m)
	}
	return fmt.Sprintf("ENV: %s", e.e, e.m)
}

type Environment struct {
	env    *types.Pair
	parent *Environment
}

func New(parent *Environment) *Environment {
	return &Environment{nil, parent}
}

func Define(env *Environment, sym *types.Symbol, item types.Type) {
	env.env = types.Cons(types.Cons(sym, item), env.env)
}

func Get(env *Environment, sym *types.Symbol) (types.Type, error) {
	for e := env; e != nil; e = e.parent {
		for v := e.env; v != nil; {
			testSym, err := types.Caar(v)
			if err != nil {
				return nil, &envError{fmt.Sprintf("Checking for %s", sym.Value), err}
			}

			if types.Eqv(sym, testSym) {
				if r, err := types.Cdar(v); err != nil {
					return nil, &envError{fmt.Sprintf("Checking for %s", sym.Value), err}
				} else {
					return r, nil
				}
			}

			nextV, _ := types.Cdr(v)

			var ok bool
			v, ok = nextV.(*types.Pair)
			if !ok {
				return nil, &envError{m: "Corrupt environment!!!"}
			}
		}
	}

	return nil, &envError{m: fmt.Sprintf("Unknown variable: %s", sym.Value)}
}

func AddArgs(env *Environment, argNames, args types.Type) (err error) {
	var nameP, valueP, name, value types.Type

	for nameP, valueP = argNames, args; nameP != nil && valueP != nil; {
		name, err = types.Car(nameP)
		if err != nil {
			return
		}

		value, err = types.Car(valueP)
		if err != nil {
			return
		}

		Define(env, name.(*types.Symbol), value)

		nameP, err = types.Cdr(nameP)
		valueP, err = types.Cdr(valueP)
	}

	if nameP != nil || valueP != nil {
		return fmt.Errorf("Wrong number of arguments. Expected %s, got %s", argNames, args)
	}
	return
}
