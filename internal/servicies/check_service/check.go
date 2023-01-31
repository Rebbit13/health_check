package check_service

import (
	"net/http"
	"strings"
)

var checksMap = map[string]healthCheck{
	"status_code": checkStatusCodeSuccess{},
	"text":        checkTextOk{},
}

type healthCheck interface {
	Pass(response *http.Response) bool
}

type checkStatusCodeSuccess struct{}

func (c checkStatusCodeSuccess) Pass(response *http.Response) bool {
	if response.StatusCode == 200 {
		return true
	}
	return false
}

type checkTextOk struct{}

func (c checkTextOk) Pass(response *http.Response) bool {
	text := make([]byte, 0)
	_, err := response.Body.Read(text)
	if err != nil {
		return false
	}
	if strings.Contains(string(text), "ok") {
		return true
	}
	return false
}
