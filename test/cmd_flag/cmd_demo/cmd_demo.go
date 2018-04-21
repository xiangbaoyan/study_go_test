package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func fileExists(fileName string) bool {
	_, e := os.Stat(fileName)
	//判断os 错误的方法
	return e == nil || os.IsExist(e)
}

func copyFile(src, dst string) (w int64, err error) {
	srcFile, e := os.Open(src)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer srcFile.Close()

	dstFile, i := os.Create(dst)
	if i != nil {
		fmt.Println(i)
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)

}

func copyFileAction(src, dst string, showProgress, force bool) {
	if !force {
		if !fileExists(dst) {
			fmt.Printf("%s exists override?y/n\n", dst)
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()
			if strings.TrimSpace(string(data)) != "y" {
				return
			}
		}
	}
	copyFile(src, dst)

	if showProgress {
		fmt.Printf("'%s'->'%s'\n", src, dst)
	}

}

func main() {
	var showProgress, force bool
	flag.BoolVar(&force, "f", false, "force copy when existing")
	flag.BoolVar(&showProgress, "v", false, "showing what is being done")

	flag.Parse()

	if flag.NArg() < 2 {
		//把用途打印出来
		flag.Usage()
		return
	}

	//fmt.Printf("fileSrc is %s,fileDst is %s",os.Args[1],os.Args[2])
	copyFileAction(os.Args[1], os.Args[2], showProgress, force)
}
