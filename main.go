package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("div by zero")
	}
	return a / b, nil
}

func main() {
	res, err := div(4, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println(res)

}
