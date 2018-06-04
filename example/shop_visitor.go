package main

import (
	"fmt"

	"github.com/go-design-patterns/behavioral/visitor"
)

func main() {
	products := make([]visitor.VisitableShop, 3)
	products[0] = &visitor.Rice{
		Product: visitor.Product{
			Price: 32.0,
			Name:  "Some rice",
		},
	}
	products[1] = &visitor.Pasta{
		Product: visitor.Product{
			Price: 40.0,
			Name:  "Some pasta",
		},
	}
	products[2] = &visitor.Fridge{
		Product: visitor.Product{
			Price: 50.0,
			Name:  "A fridge",
		},
	}

	// Print the sum of prices
	priceVisitor := &visitor.PriceVisitor{}
	for _, p := range products {
		p.Accept(priceVisitor)
	}

	fmt.Printf("Total: %0.2f\n", priceVisitor.Sum)

	// Print the product list
	nameVisitor := &visitor.NamePrinter{}
	for _, p := range products {
		p.Accept(nameVisitor)
	}

	fmt.Printf("\nProduct list:\n---------------\n%s", nameVisitor.ProductList)
}
