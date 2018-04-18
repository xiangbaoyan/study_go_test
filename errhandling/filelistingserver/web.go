package main

import (
	"fmt"
	"github.com/xiangbaoyan/study_go_test/functional/fib/filelisting"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

//返回参数类型
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	fmt.Println("进入到此")
	//这是返回函数本身
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("进入到此2")

		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound

			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
func main() {
	fmt.Println("运行中")
	//http.HandleFunc("list/errhandling/filelistingserver/web.go ")
	http.HandleFunc("/", errWrapper(filelisting.HandlerFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
