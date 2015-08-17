package types

import "fmt"

type Number struct {
	Value int64
}

func (n *Number) String() string {
	return fmt.Sprintf("%d", int64(n.Value))
}

func (n1 *Number) Eqv(n2 SchemeType) bool {
	n, ok := n2.(*Number)
	if ok {
		return n1.Value == n.Value
	} else {
		return false
	}
}

func IsNumber(n SchemeType) (ok bool) {
	_, ok = n.(*Number)
	return
}

func Add(env interface{}, args SchemeType) (SchemeType, error) {
	var total int64 = 0

	for ; args != nil; args, _ = Cdr(args){
		arg, _ := Car(args)
		n, ok := arg.(*Number)
		if !ok {
			return nil, fmt.Errorf("ADD: Not a number: %s", arg)
		}
		
		total = total + n.Value
	}
	return &Number{total}, nil
}
