package types

import "fmt"

type Symbol struct {
	Value string
}

func (s1 *Symbol) Eqv(s2 Type) bool {
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

func IsSymbol(s Type) (ok bool) {
	_, ok = s.(*Symbol)
	return
}
