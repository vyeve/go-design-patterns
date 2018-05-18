package main

import (
	"flag"
	"log"

	strategy2 "github.com/go-design-patterns/behavioral/strategy"
)

var output = flag.String("output", "console", "The output to use between"+
	" 'console' and 'image' file")

func main() {
	flag.Parse()

	strategy := strategy2.NewSquareStrategy(*output)

	if err := strategy.Print(); err != nil {
		log.Fatal(err)
	}
}
