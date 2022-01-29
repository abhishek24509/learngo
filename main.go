package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello gophers !!!")

	helloFunc := func() {
		fmt.Println("hello abhishek welcome to go world")
	}

	helloFunc()

	helloFunc3 := func() {
		fmt.Println("hello abhishek welcome to go world helloFunc3")
	}

	var helloFunc2 *func()

	helloFunc2 = &helloFunc3

	(*helloFunc2)()

}
