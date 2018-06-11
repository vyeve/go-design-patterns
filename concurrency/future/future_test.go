package future

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}

	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)

		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "Hello World!", nil
		})
		wg.Wait()
	})

	t.Run("Error result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)
		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "", errors.New("error occurred")
		})
		wg.Wait()
	})

	t.Run("Closure Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)

		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})

		future.Execute(setContext("Hello"))
		wg.Wait()
	})
}

func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("TimeOut!")

	t.Fail()
	wg.Done()
}

func setContext(msg string) ExecuteStringFunc {
	msg = fmt.Sprintf("%s Closure!\n", msg)
	return func() (string, error) {
		return msg, nil
	}
}
