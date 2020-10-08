package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	c "github.com/dmalix/financelime-rest-api-test/cmd/config"
	"io/ioutil"
	"net/http"
	"time"
)

/*
	HTTP Client
		----------------
		Return:
			statusCode   int
			err          error
*/
func HttpClient(env *c.Env, pMethod, pEndpoint string,
	pHeaders map[string]string,
	pRequestBody, pResponseBody interface{}) (int, error) {

	var (
		transport    *http.Transport
		client       *http.Client
		request      *http.Request
		response     *http.Response
		requestBody  []byte
		responseBody []byte
		statusCode   int
		err          error
	)

	transport = &http.Transport{
		IdleConnTimeout: time.Millisecond * time.Duration(env.Config.Client.IdleConnTimeoutMs),
	}

	client = &http.Client{
		Transport: transport,
		Timeout:   time.Millisecond * time.Duration(env.Config.Client.TimeoutMs),
	}

	if pRequestBody != nil {
		requestBody, err = json.Marshal(pRequestBody)
		if err != nil {
			return statusCode, errors.New(fmt.Sprintf("Failed to marshal the props to json-format [%s]", err))
		}
	}

	request, err =
		http.NewRequest(
			pMethod,
			fmt.Sprintf("%s://%s", env.Config.General.Protocol, env.Config.General.Domain+pEndpoint),
			bytes.NewBuffer(requestBody))
	if err != nil {
		return statusCode, errors.New(fmt.Sprintf("Failed to create new request [%s]", err))
	}

	for key, value := range pHeaders {
		request.Header.Set(key, value)
	}

	response, err = client.Do(request)
	if err != nil {
		return statusCode, errors.New(fmt.Sprintf("Failed to do the request [%s]", err))
	}

	//noinspection GoUnhandledErrorResult
	defer response.Body.Close()

	statusCode = response.StatusCode

	if statusCode == 200 {

		responseBody, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return statusCode, errors.New(fmt.Sprintf("Failed to read the response [%s]", err))
		}

		if pResponseBody != nil {
			err = json.Unmarshal(responseBody, &pResponseBody)
			if err != nil {
				return statusCode, errors.New(fmt.Sprintf("Failed to unmarshal the body to struct-format [%s]", err))
			}
		}
	}

	return statusCode, nil
}
