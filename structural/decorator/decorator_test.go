package decorator

import (
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}
	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "Pizza with the following ingredients:"

	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("when calling the add ingredient of the pizza decorator "+
			"it must return the text: \n%#q the expected text,\n not %#q", expectedText, pizzaResult)
	}
}

func TestOnion_AddIngredient(t *testing.T) {
	onion := &Onion{}
	onionResult, err := onion.AddIngredient()

	if err == nil {
		t.Errorf("when calling AddIngredient on the onion decorator without an IngredientAdd "+
			"on its Ingredient field must return an error, not a string with %#q", onionResult)
	}

	onion = &Onion{&PizzaDecorator{}}
	onionResult, err = onion.AddIngredient()

	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(onionResult, "onion") {
		t.Errorf("when calling the add ingredient on the onion decorator it must return"+
			"a text with the word `onion`, not %#q", onionResult)
	}
}

func TestMeat_AddIngredient(t *testing.T) {
	meat := new(Meat)
	meatResult, err := meat.AddIngredient()

	if err == nil {
		t.Errorf("when calling AddIngredient on the meat decorator without an IngredientAdd"+
			"on its Ingredient field must return an error, not a string with %#q", meatResult)
	}

	meat = &Meat{new(PizzaDecorator)}
	meatResult, err = meat.AddIngredient()

	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(meatResult, "meat") {
		t.Errorf("when calling the add ingredient on the meat decorator it must return"+
			"a text with the word `meat`, not %#q", meatResult)
	}
}

func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &Onion{&Meat{&PizzaDecorator{}}}
	pizzaResult, err := pizza.AddIngredient()

	if err != nil {
		t.Error(err)
	}

	expectedText := "Pizza with the following ingredients: meat, onion"

	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("when asking for a pizza with onion and meat the returned string must"+
			"contain the text %#q but %#q didn't have it", expectedText, pizzaResult)
	}
	t.Log(pizzaResult)
}
