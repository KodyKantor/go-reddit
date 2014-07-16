package api

import (
	"io/ioutil"
	"net/http"
)

func ProcessRequest(request string, method string) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, request, nil)
	req.Header.Set("User-Agent", AGENT)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
