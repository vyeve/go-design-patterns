package main

import (
	"fmt"

	"github.com/go-design-patterns/behavioral/memento"
)

func main() {
	m := memento.MementoFacade{}

	m.SaveSettings(memento.Volume(4))
	m.SaveSettings(memento.Mute(false))

	assertAndPrint(m.RestoreSettings(0))
	assertAndPrint(m.RestoreSettings(1))
}

func assertAndPrint(c memento.Command) {
	switch cast := c.(type) {
	case memento.Volume:
		fmt.Printf("Volume:\t%d\n", cast)
	case memento.Mute:
		fmt.Printf("Mute:\t%t\n", cast)
	}
}
