package mediator

import "errors"

type (
	One   struct{}
	Two   struct{}
	Three struct{}
	Four  struct{}
	Five  struct{}
	Six   struct{}
	Seven struct{}
	Eight struct{}
	Nine  struct{}
	Zero  struct{}
)

func (o *One) OnePlus(n interface{}) interface{} {
	switch n.(type) {
	case One:
		return &Two{}
	case Two:
		return &Three{}
	case Three:
		return &Four{}
	case Four:
		return &Five{}
	case Five:
		return &Six{}
	case Six:
		return &Seven{}
	case Seven:
		return &Eight{}
	case Eight:
		return &Nine{}
	case Nine:
		return [2]interface{}{&One{}, &Zero{}}
	default:
		return errors.New("number not found")
	}
}

func Sum(a, b interface{}) interface{} {
	switch a := a.(type) {
	case One:
		switch b := b.(type) {
		case One:
			return &Two{}
		case Two:
			return &Three{}
		case int:
			return b + 1
		default:
			return errors.New("number not found")
		}
	case Two:
		switch b := b.(type) {
		case One:
			return &Two{}
		case Two:
			return &Three{}
		case int:
			return b + 2
		default:
			return errors.New("number not found")
		}
	case int:
		switch b := b.(type) {
		case One:
			return &Two{}
		case Two:
			return &Three{}
		case int:
			return a + b
		default:
			return errors.New("number not found")
		}
	default:
		return errors.New("number not found")
	}
}
