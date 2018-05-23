package main

import (
	"fmt"
	"time"

	"github.com/go-design-patterns/behavioral/command"
)

var com command.Commander

func main() {
	com = &command.HelloMessage{}
	fmt.Println(com.Info())
	com = &command.TimePassed{Start: time.Now()}
	time.Sleep(time.Second)
	fmt.Println(com.Info())
}
