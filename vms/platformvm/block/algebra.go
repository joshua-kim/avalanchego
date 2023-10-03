package block

import "fmt"

type Algebra[T Interface] interface {
	Verify()
	Reject()
	Accept()
}

func NewAlgebra(t Interface) (Algebra[Interface], error) {
	switch t.(type) {
	case Banff:
		return nil, nil
	default:
		return nil, fmt.Errorf("unexpected type")
	}
}
