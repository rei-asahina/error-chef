package main

import (
	"errors"
	"fmt"

	"github.com/rei-asahina/error-chef/cooking"
)

func doSomething() error {
	return errors.New("実際のエラー")
}

func main() {
	err := doSomething()
	if err != nil {
		wrappedErr := cooking.Wrap(err, "doSomethingでエラー発生")
		fmt.Println("エラー:", wrappedErr)
	}
}
