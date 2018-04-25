package main

import "fmt"

func main() {

	arr := "hello"

	v, ok := interface{}(arr).(string)
	fmt.Println(v, ok)
}
