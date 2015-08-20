package types

import (
	"bytes"
	"errors"
	"fmt"
)

type Pair struct {
	car SchemeType
	cdr SchemeType
}

func Cons(a, b SchemeType) *Pair {
	return &Pair{a, b}
}

func Car(p SchemeType) (SchemeType, error) {
	pair, ok := p.(*Pair)
	if ok {
		return pair.car, nil
	}
	return nil, errors.New(fmt.Sprintf("Car: Not pair: %s", p))
}

func Cdr(p SchemeType) (SchemeType, error) {
	pair, ok := p.(*Pair)
	if ok {
		return pair.cdr, nil
	}
	return nil, errors.New(fmt.Sprintf("Cdr: Not pair: %s", p))
}

func Caar(p SchemeType) (SchemeType, error) {
	p1, err := Car(p)
	if err != nil {
		return nil, err
	}
	return Car(p1)
}

func Cddr(p SchemeType) (SchemeType, error) {
	p1, err := Cdr(p)
	if err != nil {
		return nil, err
	}
	return Cdr(p1)
}

func Cdar(p SchemeType) (SchemeType, error) {
	p1, err := Car(p)
	if err != nil {
		return nil, err
	}
	return Cdr(p1)
}

func Cadr(p SchemeType) (SchemeType, error) {
	p1, err := Cdr(p)
	if err != nil {
		return nil, err
	}
	return Car(p1)
}

func (p1 *Pair) Eqv(p2 SchemeType) bool {
	p, ok := p2.(*Pair)
	if ok {
		return p1.car == p.car && p1.cdr == p.cdr
	} else {
		return false
	}
}

func handleNull(t SchemeType) string {
	if t == nil {
		return "#nil"
	}
	return fmt.Sprintf("%s", t)
}

func (p *Pair) String() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("(%s", handleNull(p.car)))

	isPair := true

	for i := 0; i < 10 && isPair; i++ {
		if p, isPair = p.cdr.(*Pair); isPair {
			b.WriteString(fmt.Sprintf(" %s", handleNull(p.car)))
		}
	}

	if p != nil {
		b.WriteString(fmt.Sprintf(" . %s)", p))
	} else {
		b.WriteString(")")
	}

	return b.String()
}

func IsPair(p SchemeType) (ok bool) {
	_, ok = p.(*Pair)
	return
}
