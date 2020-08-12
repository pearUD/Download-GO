package response

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var client http.Client

func GetResponse(request *http.Request) *http.Response {
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("链接异常 %s", err.Error())
	}
	return response
}

func CreateRequest(url, method string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("链接异常 %s", err.Error())
	}
	return request
}

func ChangeHead(request *http.Request, k, v string) *http.Request {
	if k != "" && v != "" {
		request.Header.Set(k, v)
	}
	return request
}

func GetHead(request *http.Request, k, v string) bool {
	if request.Header.Get(k) == v {
		return true
	}
	return false
}

func Condition(response *http.Response, maps map[string]string) bool {
	for k, v := range maps {
		fmt.Println(response.Header.Get(k))
		if !(response.Header.Get(k) == v) {
			return false
		}
	}
	return true
}
