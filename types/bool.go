package types

type Bool struct {
	Value bool
}

func (b *Bool) String() string {
	if b.Value {
		return "#t"
	}
	return "#f"
}

func (b1 *Bool) Eqv(b2 Type) bool {
	b, ok := b2.(*Bool)
	if ok {
		return b1.Value == b.Value
	} else {
		return false
	}
}

func IsBool(b Type) (ok bool) {
	_, ok = b.(*Bool)
	return
}
