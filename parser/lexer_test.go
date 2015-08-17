package parser

import (
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	const str = "(moo cows (* 4 5) \"moo \\\"cows\\\"\")"
	result := []string{"(", "moo", "cows", "(", "*", "4", "5", ")",
		               "\"moo \\\"cows\\\"\"", ")"}

	tk := newTokenizer(strings.NewReader(str))

	for _, r := range(result) {
		if s, err := tk.nextToken(); s != r || err != nil {
			t.Error("Token not '", r, "': ", s)
		}
	}

	if s, err := tk.nextToken() ; err == nil {
		t.Error("Buffer overrun: ", s)
	}
}
