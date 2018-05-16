package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-design-patterns/structural/decorator"
)

func main() {
	fmt.Println("Enter the type number of server you want to launch from the following:")
	fmt.Println("1.- Plain server")
	fmt.Println("2.- Server with logging")
	fmt.Println("3.- Server with logging and authentication")

	var selection int
	fmt.Fscanf(os.Stdin, "%d", &selection)

	var mySuperServer http.Handler
	switch selection {
	case 1:
		mySuperServer = new(decorator.MyServer)
	case 2:
		mySuperServer = &decorator.LoggerServer{
			Handler:   new(decorator.MyServer),
			LogWriter: os.Stdout,
		}
	case 3:
		var user, password string
		fmt.Println("Enter user and password separated by a space")
		fmt.Fscanf(os.Stdin, "%s %s", &user, &password)
		mySuperServer = &decorator.LoggerServer{
			Handler: &decorator.BasicAuthMiddleware{
				Handler:  new(decorator.MyServer),
				User:     user,
				Password: password,
			},
			LogWriter: os.Stdout,
		}
	default:
		mySuperServer = new(decorator.MyServer)

	}

	http.Handle("/", mySuperServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
