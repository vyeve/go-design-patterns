package main

import (
	"fmt"

	"github.com/go-design-patterns/behavioral/mediator"
)

func main() {
	fmt.Printf("%#v\n", mediator.Sum(mediator.One{}, mediator.Two{}))
	fmt.Printf("%#v\n", mediator.Sum(1, 2))
	fmt.Printf("%#v\n", mediator.Sum(mediator.Two{}, 2))
}
