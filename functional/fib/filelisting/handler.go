package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HandlerFileList(writer http.ResponseWriter, request *http.Request) error {

	fmt.Println("7777")
	path := request.URL.Path[len("/list/"):]
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
	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		return err
	}
	return nil
}
