package parser

import (
	"strings"
	"testing"
	"github.com/alanthird/gscheme/types"
)

func TestParser(t *testing.T) {
	r, err := Parse(strings.NewReader("sym"))
	if err != nil {
		t.Error(err)
	} else if !types.Eqv(r, &types.Symbol{"sym"}) {
		t.Error("Expecting symbol, got", r)
	}

	str := "(moo (said the many) cows)"

	r, err = Parse(strings.NewReader(str))
	if err != nil {
		t.Error(err)
	} else if r.String() != str {
		t.Error("expected", str, "got", r)
	}

	str = "(moo (cows"

	r, err = Parse(strings.NewReader(str))
	if err == nil {
		t.Error("Expected parser error, got", r)
	}

	str = "(\"moo\" said the cows)"
	r, err = Parse(strings.NewReader(str))
	if err != nil {
		t.Error(err)
	} else if r.String() != str {
		t.Error("expected", str, "got", r)
	}

	str = "'(moo ())"
	expected := "(quote (moo #nil))"
	r, err = Parse(strings.NewReader(str))
	if err != nil {
		t.Error(err)
	} else if r.String() != expected {
		t.Error("expected", expected, "got", r)
	}

	str = "(+ 1 1)"
	expected = "(+ 1 1)"
	r, err = Parse(strings.NewReader(str))
	if err != nil {
		t.Error(err)
	} else if r.String() != expected {
		t.Error("expected", expected, "got", r)
	}
}
