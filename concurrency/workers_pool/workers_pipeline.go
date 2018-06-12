package workers_pool

import (
	"fmt"
	"log"
	"sync"
)

type Request struct {
	Data    interface{}
	handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	myRequest := Request{
		Data: fmt.Sprintf(s, id),
		handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
	return myRequest
}
