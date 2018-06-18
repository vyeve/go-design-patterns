package error_handling

import (
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func CheckStatus(done <-chan interface{}, urls ...string) <-chan Result {
	results := make(chan Result)

	go func() {
		defer close(results)

		for _, url := range urls {
			resp, err := http.Get(url)
			result := Result{
				Error:    err,
				Response: resp,
			}
			select {
			case <-done:
				return
			case results <- result:
			}
		}
	}()

	return results
}

/*
//Example:

done := make(chan interface{})
defer close(done)

urls := []string{"https://www.google.com", "https://badhost"}
for result := range CheckStatus(done, urls...) {
	if result.Error != nil {
		log.Printf("error: %v", result.Error)
		continue
	}
	fmt.Printf("Response: %v\n", result.Response.Status)
}
*/
