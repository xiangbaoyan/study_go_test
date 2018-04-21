package main

import (
	"flag"
	"fmt"
)

func main() {
	style3()
	flag.PrintDefaults()
}

func style3() {
	//新样式，直接传入变量
	var method string
	var val int
	flag.StringVar(&method, "method", "默认值", "style3模式")
	flag.IntVar(&val, "value", -1, "style3模式")
	flag.Parse()
	fmt.Println(method, val)

}

func style2() {
	//格式化定义
	methodPtr := flag.String("method", "default", "method of sample")
	value := flag.Int("value", -1, "value of sample")
	//解析
	flag.Parse()
	fmt.Println(*methodPtr, *value)
}
