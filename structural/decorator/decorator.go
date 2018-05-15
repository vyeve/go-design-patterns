/*
The Decorator design pattern allows you to decorate an already existing type with more
functional features without actually touching it. It uses an approach similar to matryoshka
doll.
The Decorator type implements the same interface of the type it decorates, and stores an instance
of that type in its members.
*/

package decorator

import (
	"errors"
	"fmt"
)

type IngredientAdd interface {
	AddIngredient() (string, error)
}

type PizzaDecorator struct {
	Ingredient IngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

type (
	Meat struct {
		Ingredient IngredientAdd
	}
	Onion struct {
		Ingredient IngredientAdd
	}
)

func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("an IngredientAdd is needed in the Ingredient field of the Meat")
	}

	s, err := m.Ingredient.AddIngredient()

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "meat"), nil
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("an IngredientAdd is needed in the Ingredient field of the Onion")
	}

	s, err := o.Ingredient.AddIngredient()

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", s, "onion"), nil
}
