package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	file, e := os.Open("avatar.png")
	if e != nil {
		panic(e)
	}
	var hA, hB byte

	binary.Read(file, binary.BigEndian, &hA)
	binary.Read(file, binary.BigEndian, &hB)

	var size uint32
	binary.Read(file, binary.BigEndian, &size)
	var r1, r2 byte
	binary.Read(file, binary.BigEndian, &r1)
	binary.Read(file, binary.BigEndian, &r2)

	fmt.Println(hA, hB, size, r1, r2)

}
