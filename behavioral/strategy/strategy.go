/*
The Strategy pattern uses different algorithms to achieve some specific functionality. These
algorithms are hidden behind an interface and, of course, they must be interchangeable. All
algorithms achieve the same functionality in a different way. For example, we could have a
`Sort` interface and few sorting algorithms. The result is the same, some list of sorted, but we
could have used quick sort, merge sort, and so on.
*/

package strategy

type PrintStrategy interface {
	Print() error
}

func NewSquareStrategy(strategy string) PrintStrategy {
	switch strategy {
	case "console":
		return &ConsoleSquare{}
	case "image":
		return &ImageSquare{DestinationFilePath: "/tmp/square.jpg"}
	default:
		return &ConsoleSquare{}
	}
}
