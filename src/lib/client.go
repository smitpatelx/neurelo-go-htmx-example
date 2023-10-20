package lib

import (
	"fmt"
	"net/http"
	"os"
)

var ApiClient *http.Client

func SetupClient() {
	ApiClient = &http.Client{}
}

func GetRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://us-east-2.aws.neurelo.app%s", url), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", os.Getenv("NEURELO_API_KEY"))
	return req, nil
}

func Call(url string) (*http.Response, error) {
	req, err := GetRequest(url)
	if err != nil {
		return nil, err
	}
	return ApiClient.Do(req)
}
