package environment

import (
	"github.com/alanthird/gscheme/types"
	"testing"
)

func TestEnv(t *testing.T) {
	e1 := New(nil)

	Define(e1, &types.Symbol{"cows"}, &types.String{"moo"})
	Define(e1, &types.Symbol{"sheep"}, &types.String{"baa"})

	e2 := New(e1)

	Define(e2, &types.Symbol{"chickens"}, &types.String{"bawk"})
	Define(e2, &types.Symbol{"sheep"}, &types.String{"meh"})

	if r, _ := Get(e2, &types.Symbol{"sheep"}); !types.Eqv(r, &types.String{"meh"}) {
		t.Error("Sheep in e2 do not say 'meh'")
	}

	if r, _ := Get(e1, &types.Symbol{"sheep"}); !types.Eqv(r, &types.String{"baa"}) {
		t.Error("Sheep in e1 do not say 'baa': ", r.(*types.String).Value)
	}

	Define(e2, &types.Symbol{"sheep"}, &types.String{"baa"})

	if r, _ := Get(e2, &types.Symbol{"sheep"}); !types.Eqv(r, &types.String{"baa"}) {
		t.Error("Sheep in e2 do not say 'baa': ", r.(*types.String).Value)
	}
}
