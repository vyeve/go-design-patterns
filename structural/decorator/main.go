package decorator

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", new(MyServer))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
