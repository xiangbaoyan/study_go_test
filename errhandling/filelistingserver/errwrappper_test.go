package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}
func (e testingUserError) Message() string {
	return string(e)
}

func errPanic(writer http.ResponseWriter, r *http.Request) error {
	panic("123")
}
func errUserErr(writer http.ResponseWriter, r *http.Request) error {
	return testingUserError("user error")
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	//{errPanic,500,"Internal Server Error"},
	{errUserErr, 400, "user error"},
}

func TestErrWrapper(t *testing.T) {
	//Internal Server Error

	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com",
			nil,
		)
		f(response, request)
		verifyResponse(response.Result(), tt.code, tt.message, t)
		//b, _ := ioutil.ReadAll(response.Body)
		//body := strings.Trim(string(b), "\n")
		//if response.Code != tt.code ||
		//	body != tt.message {
		//	t.Errorf("expected (%d,%s);got(%d,%s)", tt.code, tt.message, response.Code, body)
		//}
	}

}

//真调用
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMessage string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode ||
		body != expectedMessage {
		t.Errorf("expected (%d,%s);got(%d,%s)", expectedCode, expectedMessage, resp.StatusCode, body)
	}
}
