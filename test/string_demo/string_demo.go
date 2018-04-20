package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ss := "dd#d444#999#"
	split := strings.Split(ss, "#")

	fmt.Println(strings.Join(split, "$"))

	fmt.Println(strconv.Atoi("99#9"))
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseFloat("3.153", 32))

	//第二个参数为进制
	fmt.Println(strconv.FormatInt(123, 2))
}
