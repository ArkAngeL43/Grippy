package GRIP

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Req_with_Response(uri string) (string, *http.Response, error) {
	var HTTP_ = http.Client{}
	request, X := http.NewRequest("GET", uri, nil)
	if X != nil {
		fmt.Printf("\033[38;5;55m|\033[38;5;88m-\033[38;5;55m| \033[38;5;88m Got a error when trying to make a request to the url, you might be offline, or connection may be private or unstable, please try again following this error -> %s ", X)
		return "", nil, X
	}
	response, X := HTTP_.Do(request)
	if X != nil {
		fmt.Printf("\033[38;5;55m|\033[38;5;88m-\033[38;5;55m| \033[38;5;88m Got a error when trying to make a request to the url, you might be offline, or connection may be private or unstable, please try again following this error -> %s ", X)
		return "", nil, X
	}
	defer response.Body.Close()
	b, X := ioutil.ReadAll(response.Body)
	if X != nil {
		fmt.Printf("\033[38;5;55m|\033[38;5;88m-\033[38;5;55m| \033[38;5;88m Got a error when trying to read the response body, maybe the body is corrupted, of wrong structure, of wrong type, or just messed up from a broken connection -> %s ", X)
		return "", nil, X
	}
	return string(b), response, X
}
