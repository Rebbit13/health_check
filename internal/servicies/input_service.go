package servicies

import (
	"encoding/json"
	"health_check/internal/domain"
	"os"
)

type JsonConfigFileInput struct {
	path string
}

func NewJsonConfigFileInput(path string) *JsonConfigFileInput {
	return &JsonConfigFileInput{path: path}
}

func (j *JsonConfigFileInput) parseJson(data []byte) ([]*domain.SiteToCheck, error) {
	type Config struct {
		Urls []*domain.SiteToCheck `json:"urls"`
	}
	config := Config{
		Urls: make([]*domain.SiteToCheck, 0),
	}
	err := json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config.Urls, nil
}

func (j *JsonConfigFileInput) GetSitesToCheck() ([]*domain.SiteToCheck, error) {
	data, err := os.ReadFile(j.path)
	if err != nil {
		return nil, err
	}
	return j.parseJson(data)
}
