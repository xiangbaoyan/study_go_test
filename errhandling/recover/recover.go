package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error Occurred:", err)
		} else {
			panic(r)
		}
	}()

	b := 0
	a := 5 / b

	fmt.Println(a)
	//panic(errors.New("this is an  Error"))
}

func main() {

	tryRecover()
}
