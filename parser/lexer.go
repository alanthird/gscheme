package parser

import (
	"io"
	"bufio"
	"bytes"
)

type tokenizer struct {
	r *bufio.Reader
}

func newTokenizer(r io.Reader) *tokenizer {
	return &tokenizer{bufio.NewReader(r)}
}

func (t *tokenizer) read() (rune, error) {
	c, _, err := t.r.ReadRune()
	if err != nil {
		return rune(0), err
	}

	return c, nil
}

func (t *tokenizer) rewind() {
	t.r.UnreadRune()
}

func (t *tokenizer) readWhitespace() {
	for c, err := t.read() ; isWhitespace(c) ; {
		if err != nil {
			return
		}
		c, err = t.read()
	}
	t.rewind()
}

func (t *tokenizer) readString() string {
	var b bytes.Buffer
	var c rune
	var err error
	
	for c, err = t.read() ; err == nil && !isSeparator(c) ; c, err = t.read() {
		b.WriteRune(c)
	}

	if err == nil {
		t.rewind()
	}

	return b.String()
}

func (t *tokenizer) readQuotedString() string {
	var b bytes.Buffer

	// the first rune *must* be a double quote or we wouldn't be here
	c, _ := t.read()
	b.WriteRune(c)
	
	for {
		c, err := t.read()
		if err != nil {
			return b.String()
		}

		b.WriteRune(c)

		if isEscapeChar(c) {
			c, err := t.read()
			if err != nil {
				return b.String()
			}
			b.WriteRune(c)
		} else if isDoubleQuote(c) {
			return b.String()
		}
	}
}

func (t *tokenizer) nextToken() (string, error) {
	t.readWhitespace()

	c, err := t.read()
	if err != nil {
		return "", err
	}
	if isParen(c) {
		return string(c), nil
	}
	if isDoubleQuote(c) {
		t.rewind()
		return t.readQuotedString(), nil
	}
	if isSingleQuote(c) {
		return string(c), nil
	}

	t.rewind()
	return t.readString(), nil
}
