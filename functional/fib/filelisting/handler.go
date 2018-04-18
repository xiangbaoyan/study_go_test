package filelisting

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}
func (e userError) Message() string {
	return string(e)
}

const prefix = "/list/"

func HandlerFileList(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println("草您个孙子1")

	if strings.Index(request.URL.Path, prefix) != 0 {
		fmt.Println("草您个孙子3")
		//panic("Path must start with "+prefix)
		return userError("Path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
