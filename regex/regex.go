package main

import (
	"fmt"
	"regexp"
)

const text = `My email is 271190187@qq.com

email is aaad@qq.com
email is bbb@qq.com
email is cccc@qq.com
email is dddd@qq.com`

func main() {

	reg := `([a-zA-Z0-9]+)@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`
	compile := regexp.MustCompile(reg)
	match := compile.FindAllStringSubmatch(text, 2)
	fmt.Println(match)
}
