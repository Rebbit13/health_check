package check_service

import (
	"fmt"
	"health_check/internal/domain"
	"net/http"
	"strings"
)

var checksMap = map[string]healthCheck{
	"status_code": checkStatusCodeSuccess{},
	"text":        checkTextOk{},
}

type healthCheck interface {
	Pass(response *http.Response) domain.Check
}

// default check if response got
type connection struct{}

func (c connection) Pass(response *http.Response) domain.Check {
	if response == nil {
		return domain.Check{
			Title:          "connection",
			Pass:           false,
			ExpectedResult: []byte("got response"),
			GotResult:      []byte("have no response"),
		}
	}
	return domain.Check{
		Title:          "connection",
		Pass:           true,
		ExpectedResult: []byte("got response"),
		GotResult:      []byte("got response"),
	}
}

type checkStatusCodeSuccess struct{}

func (c checkStatusCodeSuccess) Pass(response *http.Response) domain.Check {
	check := false
	if response.StatusCode == 200 {
		check = true
	}
	return domain.Check{
		Title:          "status_code",
		Pass:           check,
		ExpectedResult: []byte("200"),
		GotResult:      []byte(fmt.Sprintf("%d", response.StatusCode)),
	}
}

type checkTextOk struct{}

func (c checkTextOk) Pass(response *http.Response) domain.Check {
	text := make([]byte, 0)
	_, err := response.Body.Read(text)
	if err != nil || !strings.Contains(string(text), "ok") {
		return domain.Check{
			Title:          "text",
			Pass:           false,
			ExpectedResult: []byte("ok"),
			GotResult:      []byte(""),
		}
	}
	return domain.Check{
		Title:          "text",
		Pass:           true,
		ExpectedResult: []byte("ok"),
		GotResult:      []byte("ok"),
	}
}
