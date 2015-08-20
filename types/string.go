package types

import "fmt"

type String struct {
	Value string
}

func (s *String) String() string {
	return fmt.Sprintf("\"%s\"", string(s.Value))
}

func (s1 *String) Eqv(s2 Type) bool {
	s, ok := s2.(*String)
	if ok {
		return s1.Value == s.Value
	} else {
		return false
	}
}

func IsString(s Type) (ok bool) {
	_, ok = s.(*String)
	return
}
