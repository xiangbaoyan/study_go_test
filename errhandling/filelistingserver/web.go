package main

import (
	"github.com/xiangbaoyan/study_go_test/functional/fib/filelisting"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

//返回参数类型
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	//这是返回函数本身
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
func main() {
	//http.HandleFunc("list/errhandling/filelistingserver/web.go ")
	http.HandleFunc("/list/", errWrapper(filelisting.HandlerFileList))
}
