package parser

func isWhitespace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

func isOpeningParen(c rune) bool {
	return c == '('
}

func isClosingParen(c rune) bool {
	return c == ')'
}

func isParen(c rune) bool {
	return isOpeningParen(c) || isClosingParen(c)
}

func isSeparator(c rune) bool {
	return isParen(c) || isWhitespace(c)
}

func isNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

func isEscapeChar(c rune) bool {
	return c == '\\'
}

func isSingleQuote(c rune) bool {
	return c == '\''
}

func isDoubleQuote(c rune) bool {
	return c == '"'
}

func isFullStop(c rune) bool {
	return c == '.'
}
