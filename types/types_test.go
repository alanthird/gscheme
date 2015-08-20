package types

import (
	"testing"
)

func TestPair(t *testing.T) {
	p := Cons(nil, nil)

	if r, _ := Car(p); r != nil {
		t.Error("Car doesn't return car")
	}

	if r, _ := Cdr(p); r != nil {
		t.Error("Cdr doesn't return cdr")
	}

	q := Cons(nil, nil)

	if !p.Eqv(q) {
		t.Error("Pairs with null car and cdr not eqv")
	}

	p = Cons(Cons(nil, nil), Cons(nil, nil))
	q = Cons(Cons(nil, nil), Cons(nil, nil))

	if Eqv(p, q) {
		t.Error("Pairs with non-eq contents eqv")
	}

	car, _ := Car(p)
	cdr, _ := Cdr(p)
	
	q = Cons(car, cdr)

	if !Eqv(p, q) {
		t.Error("Pairs with eq contents not eqv")
	}
	
	if !IsPair(p) {
		t.Error("isPair gives false with Pair")
	}

	if IsPair(&String{"Moo!"}) {
		t.Error("isPair gives true with String")
	}
}

func TestNumber(t *testing.T) {
	var n Number
	n = Number{0}

	if n.Value != 0 {
		t.Error("Value != 0")
	}

	if !Eqv(&n, &Number{0}) {
		t.Error("Numbers with eq Values not eqv")
	}

	if Eqv(&n, &Number{1}) {
		t.Error("Numbers with !eq Values eqv")
	}

	if !IsNumber(&n) {
		t.Error("isNumber gives false with Number")
	}

	if IsNumber(&String{"Moo!"}) {
		t.Error("isNumber gives true with String")
	}
}

func TestString(t *testing.T) {
	var s String

	s = String{"Hello, world!"}

	if s.Value != "Hello, world!" {
		t.Error("Value != 'Hello, World!'")
	}

	if !Eqv(&s, &String{"Hello, world!"}) {
		t.Error("Strings with eq Values not eqv")
	}

	if Eqv(&s, &String{"What's up?"}) {
		t.Error("Strings with !eq Values eqv")
	}

	if !IsString(&s) {
		t.Error("isString gives false with String")
	}

	if IsString(&Number{1}) {
		t.Error("isString gives true with Number")
	}
}

func TestEq(t *testing.T) {
	var a, b *Pair

	if Eq(a, b) != true {
		t.Error("Null pairs not eq")
	}

	a = &Pair{}
	b = &Pair{}
	if Eq(a, b) != false {
		t.Error("Different pairs are eq")
	}

	b = a
	if Eq(a, b) != true {
		t.Error("Identical pairs are not eq")
	}
}

func TestEqv(t *testing.T) {
	if Eqv(&Number{42}, &String{"forty two"}) {
		t.Error("Number == String")
	}
}
