package strategy

import "fmt"

type ConsoleSquare struct{}

func (*ConsoleSquare) Print() error {
	_, err := fmt.Println("Square")
	return err
}
