package service

import (
	"errors"
	"fmt"
	c "github.com/dmalix/financelime-rest-api-test/cmd/config"
	u "github.com/dmalix/financelime-rest-api-test/cmd/utils"
	"net/http"
	"strconv"
)

func getCurrentVersion_400_NoHeaderRequestID(env *c.Env) error {

	const (
		method             = http.MethodGet
		endpoint           = "/dist"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		dist       distData
		statusCode int
	)

	headers = make(map[string]string)

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, &dist)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func getCurrentVersion_400_InvalidRequestID(env *c.Env) error {

	const (
		method             = http.MethodGet
		endpoint           = "/dist"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		dist       distData
		statusCode int
	)

	headers = make(map[string]string)

	headers["request-id"] = "1234"

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, &dist)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func getCurrentVersion_200(env *c.Env) error {

	const (
		method             = http.MethodGet
		endpoint           = "/dist"
		statusCodeExpected = 200
	)

	var (
		err        error
		headers    map[string]string
		dist       distData
		statusCode int
	)

	headers = make(map[string]string)

	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, &dist)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}
	if len(dist.Version) == 0 || len(dist.Build) == 0 {
		return errors.New(fmt.Sprintf("REST-API service returned empty data value."))
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func getCurrentState_400_NoHeaderRequestID(env *c.Env) error {

	const (
		method             = http.MethodGet
		endpoint           = "/status"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		status     status
		statusCode int
	)

	headers = make(map[string]string)

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, &status)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func getCurrentState_400_InvalidHeaderRequestID(env *c.Env) error {

	const (
		method             = http.MethodGet
		endpoint           = "/status"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		status     status
		statusCode int
	)

	headers = make(map[string]string)

	headers["request-id"] = "1234"

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, &status)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func getCurrentState_200(env *c.Env) error {

	const (
		method              = http.MethodGet
		endpoint            = "/status"
		statusCodeExpected  = 200
		statusValueExpected = true
	)

	var (
		err        error
		headers    map[string]string
		status     status
		statusCode int
	)

	headers = make(map[string]string)

	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, &status)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	if status.API != statusValueExpected || status.DB != statusValueExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong data value: got %v-%v want %v-%v",
			status.API, status.DB, statusValueExpected, statusValueExpected))
	}

	return nil
}
