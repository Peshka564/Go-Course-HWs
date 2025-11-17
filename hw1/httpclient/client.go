package httpclient

import (
	"bufio"
	"encoding/json"
	"net/http"

	"github.com/Peshka564/Go-Course-HWs/hw1/errors"
)

const baseUrl string = "https://api.github.com"

type ErrorResponse struct {
	Message string
}

type HttpClient struct {
	apiToken string
}

func (client *HttpClient) Init(apiToken string) {
	client.apiToken = apiToken
}

func (client *HttpClient) Get(endpoint string, resObj interface{}) error {
	req, err := http.NewRequest("GET", baseUrl+endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Add("Authorization", "Bearer "+client.apiToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	data, err := ReadBody(res)
	if err != nil {
		return err
	}

	// Check if we have an error from the server
	if res.StatusCode >= 400 {
		var jsonRes ErrorResponse
		jsonErr := json.Unmarshal(data, &jsonRes)
		if jsonErr != nil {
			return jsonErr
		}
		return errors.InvalidHTTPResponse{Url: req.URL.String(), StatusCode: res.StatusCode, Message: jsonRes.Message}
	}

	jsonErr := json.Unmarshal(data, resObj)
	return jsonErr
}

func ReadBody(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)

	var data []byte = nil
	for scanner.Scan() {
		data = append(data, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
