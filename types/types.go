package types

type SchemeType interface {
	Eqv(SchemeType) bool
	String() string
}

func Eq(a, b SchemeType) bool {
	return a == b
}

func Eqv(a, b SchemeType) bool {
	return a.Eqv(b)
}
