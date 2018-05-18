package strategy

import "testing"

func ExampleConsoleSquare_Print() {
	console := new(ConsoleSquare)
	console.Print()
	// Output: Square
}

func TestNewSquareStrategy(t *testing.T) {
	strategy := NewSquareStrategy("console")
	if _, ok := strategy.(*ConsoleSquare); !ok {
		t.Error("unexpectable result, must be *ConsoleSquare")
	}

	strategy = NewSquareStrategy("image")
	if _, ok := strategy.(*ImageSquare); !ok {
		t.Error("unexpectable result, must be *ImageSquare")
	}
}
