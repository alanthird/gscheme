package parser

import (
	"fmt"
	"github.com/alanthird/gscheme/types"
	"io"
	"strconv"
	"unicode/utf8"
)

type parseError struct {
	message string
	token   string
}

func (e *parseError) Error() string {
	return fmt.Sprintf("PARSE ERROR: %s: %s", e.message, e.token)
}

func list(t *tokenizer) (types.SchemeType, error) {
	car, listEnd, err := parseToken(t)
	if err != nil {
		return nil, err
	}
	if listEnd {
		return nil, nil
	}

	cdr, err := list(t)
	if err != nil {
		return nil, err
	}

	return types.Cons(car, cdr), nil
}

func parseToken(t *tokenizer) (types.SchemeType, bool, error) {
	var token string
	var err error

	if token, err = t.nextToken(); err != nil {
		return nil, false, err
	}

	if isQuote(token) {
		cadr, _, err := parseToken(t)
		if err != nil {
			return nil, false, err
		}
		return types.Cons(&types.Symbol{"quote"}, types.Cons(cadr, nil)), false, nil
	}

	if isListEnd(token) {
		return nil, true, nil
	}
	if isList(token) {
		r, err := list(t)
		return r, false, err
	}

	if isSymbol(token) {
		return &types.Symbol{token}, false, nil
	}

	if isString(token) {
		s, err := makeString(token)
		if err != nil {
			return nil, false, err
		}
		return s, false, nil
	}

	if isNum(token) {
		n, err := strconv.ParseInt(token, 10, 64)
		if err != nil {
			return nil, false, err
		}
		return &types.Number{n}, false, nil
	}

	switch token {
	case "#t":
		return &types.Bool{true}, false, nil
	case "#f":
		return &types.Bool{false}, false, nil
	}

	return nil, false, &parseError{"Unknown type", token}
}

func Parse(t io.Reader) (types.SchemeType, error) {
	r, _, err := parseToken(newTokenizer(t))
	return r, err
}

func isSymbol(t string) bool {
	c, _ := utf8.DecodeRuneInString(t)
	return !isParen(c) && !isNumber(c) && !isDoubleQuote(c) && !isHash(c)
}

func isString(t string) bool {
	c, _ := utf8.DecodeRuneInString(t)
	return isDoubleQuote(c)
}

func isList(t string) bool {
	c, _ := utf8.DecodeRuneInString(t)
	return isOpeningParen(c)
}

func isListEnd(t string) bool {
	c, _ := utf8.DecodeRuneInString(t)
	return isClosingParen(c)
}

func isQuote(t string) bool {
	c, _ := utf8.DecodeRuneInString(t)
	return isSingleQuote(c)
}

func isNum(t string) bool {
	c, _ := utf8.DecodeRuneInString(t)
	return isNumber(c)
}

func makeString(token string) (*types.String, error) {
	lastRune, _ := utf8.DecodeLastRuneInString(token)
	if !isDoubleQuote(lastRune) {
		return nil, &parseError{"Unterminated string", token}
	}

	return &types.String{token[1 : len(token)-1]}, nil
}
