package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		//中文解析
		e := determineEncoding(resp.Body)
		utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
		all, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			panic(err)
		}
		printCityList(all)
	}
}


func printCityList(contents []byte) {
	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9A-Za-z]+)"[^>]*>([^<]+)</a>`)
	matches := compile.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City: %s,URL: %s\n", m[2], m[1])
	}

	fmt.Printf("Matched found: %d\n", len(matches))
}
