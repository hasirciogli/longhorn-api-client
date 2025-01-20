package utils

import (
	"bytes"
	"net/http"

	"github.com/hasirciogli/longhorn-api-client/models"
)

type RequesterError struct {
	Message    string
	StatusCode int
}

func Requester(method string, path string, body []byte, sConfig *models.SConfig) (*http.Response, *RequesterError) {
	req, err := http.NewRequest(method, sConfig.LonghornUiEndpoint+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, &RequesterError{Message: err.Error(), StatusCode: 500}
	}

	req.Header.Set("Content-Type", "application/json")
	if sConfig.BasicAuth != nil {
		req.Header.Set("Authorization", "Basic "+*sConfig.BasicAuth)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, &RequesterError{Message: err.Error(), StatusCode: 500}
	}

	return resp, nil
}
