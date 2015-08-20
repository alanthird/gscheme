package types

import (
	"fmt"
)

type SFunction struct {
	Function Type
	Args     *Pair
	Env      interface{}
}

func (f *SFunction) String() string {
	return fmt.Sprintf("#func{%p}", f)
}

func (f1 *SFunction) Eqv(f2 Type) bool {
	return Eq(f1, f2)
}

func IsSFunction(f Type) (ok bool) {
	_, ok = f.(*SFunction)
	return
}

type Builtin struct {
	Fn func(interface{}, Type) (Type, error)
}

func (b *Builtin) String() string {
	return fmt.Sprintf("#builtin{%p}", b)
}

func (b1 *Builtin) Eqv(b2 Type) bool {
	bTwo, ok := b2.(*Builtin)
	if ok {
		return &b1.Fn == &bTwo.Fn
	}
	return false
}

func IsBuiltin(f Type) (ok bool) {
	_, ok = f.(*Builtin)
	return
}
