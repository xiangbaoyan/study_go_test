package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
)

type body struct {
	height int
	weight int
}

type person struct {
	Name string `xml:"name,attr"`
	Age  int    `xml:"age"`
	Body body
}

func main() {

	//属性和别名的表现形式
	p := person{
		Name: "小李",
		Age:  15,
		Body: body{height: 175, weight: 180},
	}

	//bytes, e := xml.Marshal(p)
	bytes, e := xml.MarshalIndent(p, "", " ")
	fmt.Println(string(bytes))
	if e != nil {
		panic(e)
	}
	file, i := os.Create("testMarshal.txt")
	defer file.Close()
	if i != nil {
		panic(i)
	}
	writer := bufio.NewWriter(file)

	writer.Write(bytes)
	writer.Flush()

	//
	//new_p := new(person)
	//xml.Unmarshal(bytes,new_p)
	//fmt.Println(new_p)

}
