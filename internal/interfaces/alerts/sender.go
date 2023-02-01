package alerts

import (
	"bytes"
	"encoding/json"
	"health_check/internal/servicies/output_service"
	"net/http"
)

const url = "http://httpbin.org/post"

type MockSender struct{}

func (m MockSender) SendStatusChanged(changes []*output_service.StatusChanges) error {
	data, err := json.Marshal(changes)
	dataReader := bytes.NewReader(data)
	if err != nil {
		return err
	}
	_, err = http.Post(
		url,
		"application/json",
		dataReader,
	)
	return err
}
