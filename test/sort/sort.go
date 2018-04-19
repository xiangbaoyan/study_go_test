package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 43, 3, 4, 1, 6, 7, 3}
	sort.Ints(arr)
	fmt.Println(arr)

}
