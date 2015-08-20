package types

type Type interface {
	Eqv(Type) bool
	String() string
}

func Eq(a, b Type) bool {
	return a == b
}

func Eqv(a, b Type) bool {
	return a.Eqv(b)
}
