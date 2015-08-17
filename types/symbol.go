package types

import "fmt"

type Symbol struct {
	Value string
}

func (s1 *Symbol) Eqv(s2 SchemeType) bool {
	s, ok := s2.(*Symbol)
	if ok {
		return s1.Value == s.Value
	} else {
		return false
	}
}

func (s *Symbol) String() string {
	return fmt.Sprintf("%s", string(s.Value))
}

func IsSymbol(s SchemeType) (ok bool) {
	_, ok = s.(*Symbol)
	return
}
