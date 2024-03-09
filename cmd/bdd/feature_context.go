package bdd

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const requestTimeout = 30 * time.Second

type FeatureContext struct {
	apiBasePath      string
	responseBody     []byte
	statusCode       int
	boxItemsResponse boxItemsResponse
}

func NewFeatureContext(apiBasePath string) *FeatureContext {
	return &FeatureContext{
		apiBasePath: apiBasePath,
	}
}

func (f *FeatureContext) reset() {
	f.responseBody = nil
	f.statusCode = 0
	f.boxItemsResponse = boxItemsResponse{}
}

func (f *FeatureContext) sendRequest(method, endpoint string, data []byte) error {
	var err error
	var req *http.Request

	client := &http.Client{Timeout: requestTimeout}

	url := f.apiBasePath + endpoint
	var reqBody io.Reader
	if len(data) > 0 {
		reqBody = bytes.NewReader(data)
	}

	req, err = http.NewRequest(method, url, reqBody)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	f.statusCode = resp.StatusCode

	f.responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	var boxItemsResponse boxItemsResponse
	if err := json.Unmarshal(f.responseBody, &boxItemsResponse); err != nil {
		return err
	}

	f.boxItemsResponse = boxItemsResponse

	return nil
}
