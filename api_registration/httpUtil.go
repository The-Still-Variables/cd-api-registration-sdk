package api_registration

import (
	"bytes"
	"crypto/tls"

	"io/ioutil"
	"net/http"
	"time"
)

// HTTPRequest ... http request
func HTTPRequest(action string, url string, jsonStr []byte, headers []string) (string, int, http.Header) {
	// LogInfo("HTTP request to : " + action + " : " + url)
	// LogInfo("Request body : " + string(jsonStr))
	retData := ""

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   time.Second * time.Duration(30),
		Transport: tr}

	req, err := http.NewRequest(action, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if len(headers)%2 == 0 && len(headers) != 0 && headers[0] != "" {
		for i := 0; i <= len(headers)/2; i = i + 2 {
			req.Header.Add(headers[i], headers[i+1])
		}
	}

	response, err := client.Do(req)
	if err != nil {
		// LogError(err)
		return "", 0, nil
	}

	data, _ := ioutil.ReadAll(response.Body)
	retData = string(data)
	// LogInfo("Response : " + retData)
	defer response.Body.Close()
	return retData, response.StatusCode, response.Header
}
