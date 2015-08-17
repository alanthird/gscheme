package types

import (
	"fmt"
)

type SFunction struct {
	Function SchemeType
	Args *Pair
	Env interface{}
}

func (f *SFunction) String() string {
	return fmt.Sprintf("#func{%p}", f)
}

func (f1 *SFunction) Eqv(f2 SchemeType) bool {
	return Eq(f1, f2)
}

func IsSFunction(f SchemeType) (ok bool) {
	_, ok = f.(*SFunction)
	return
}



type Builtin struct {
	Fn func(interface{}, SchemeType) (SchemeType, error)
}

func (b *Builtin) String() string {
	return fmt.Sprintf("#builtin{%p}", b)
}

func (b1 *Builtin) Eqv(b2 SchemeType) bool {
	bTwo, ok := b2.(*Builtin)
	if ok {
		return &b1.Fn == &bTwo.Fn
	}
	return false
}

func IsBuiltin(f SchemeType) (ok bool) {
	_, ok = f.(*Builtin)
	return
}
