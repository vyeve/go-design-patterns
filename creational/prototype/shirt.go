/*
The aim of Prototype pattern is to have an object or a set of objects that is already created at compilation time,
but which you can clone as many times as you want at runtime.The key difference between this and a Builder pattern
is that objects are cloned for the user instead of building them at runtime.
*/

package prototype

import (
	"errors"
	"fmt"
)

type ShirtCloner interface {
	GetClone(colorType ShirtColor, sku string) (ItemInfoGetter, error)
}

type ShirtColor byte

const (
	White ShirtColor = iota
	Black
	Blue
)

func GetShirtsCloner() ShirtCloner {
	return &ShirtCache{}
}

type ShirtCache struct{}

func (s *ShirtCache) GetClone(color ShirtColor, sku string) (ItemInfoGetter, error) {
	var newItem = new(Shirt)
	switch color {
	case White:
		*newItem = *whitePrototype
	case Black:
		*newItem = *blackPrototype
	case Blue:
		*newItem = *bluePrototype
	default:
		return nil, errors.New("shirt model not recognized")
	}
	newItem.SKU = sku
	return newItem, nil
}

type ItemInfoGetter interface {
	GetInfo() string
	GetPrice() float32
}

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %.2f\n",
		s.SKU, s.Color, s.Price)
}

var whitePrototype = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}
var bluePrototype = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}
