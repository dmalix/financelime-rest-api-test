package authorization

import (
	"errors"
	"fmt"
	c "github.com/dmalix/financelime-functional-tests/cmd/config"
	g "github.com/dmalix/financelime-functional-tests/cmd/pgraphics"
	u "github.com/dmalix/financelime-functional-tests/cmd/utils"
	"net/http"
	"strconv"
)

func (checkList *checkLists) revokeRefreshToken(env *c.Env, accessToken, refreshToken, publicSessionID string) (int, error) {

	const checkListName = "Revoke Refresh Token"
	const indentBeforeStatus = "\t\t\t\t\t\t\t"
	const errorMessage = "Failed to complete the %s test %s [%s]"
	var err error
	var tests = 0
	var testID string

	fmt.Print(g.ItemII)
	fmt.Print(checkListName, indentBeforeStatus)

	testID = "#M678F5tr" // No the RequestID header
	tests++
	err = test.revokeRefreshToken(env,
		accessToken,
		publicSessionID,
		http.MethodDelete,
		"/v1/oauth/sessions",
		"application/json;charset=utf-8",
		"",
		400)
	if err != nil {
		return tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}

	testID = "#2NKyviy2" // No the ContentType header
	tests++
	err = test.revokeRefreshToken(env,
		accessToken,
		publicSessionID,
		http.MethodDelete,
		"/v1/oauth/sessions",
		"",
		"K7800-H7625-Z5852-N1693-K1972",
		400)
	if err != nil {
		return tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}

	testID = "#iy2NKyv2" // Invalid Access Token
	tests++
	err = test.revokeRefreshToken(env,
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJGaW5hbmNlbGltZS5jb20iLCJzdWIiOiJBdXRob3JpemF0aW9uIiwicHVycG9zZSI6ImFjY2VzcyIsImlkIjoiMTZkN2RiNTM3MjQ3ZWFmMTEzZjRjOGFkNTllOWEyYTU4OWNlN2NhZjYxMzViY2Q3YmZlYzBiNTI1YWY0OGEyYSIsImlhdCI6MTU5NjgzNTM5OX0.d68bea3232f10c60483f838fff8d8c66661cb497b141213c9a006be2e7c9d723",
		publicSessionID,
		http.MethodDelete,
		"/v1/oauth/sessions",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		401)
	if err != nil {
		return tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}

	testID = "#Wp4oRz1n" // Invalid the PUT method
	tests++
	err = test.revokeRefreshToken(env,
		accessToken,
		publicSessionID,
		http.MethodPut,
		"/v1/oauth/sessions",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		405)
	if err != nil {
		return tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}

	testID = "#wz0FKQHO" // Invalid the POST method
	tests++
	err = test.revokeRefreshToken(env,
		accessToken,
		publicSessionID,
		http.MethodPost,
		"/v1/oauth/sessions",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		405)
	if err != nil {
		return tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}

	testID = "#Gq7dHBlM" // OK
	tests++
	err = test.revokeRefreshToken(env,
		accessToken,
		publicSessionID,
		http.MethodDelete,
		"/v1/oauth/sessions",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		204)
	if err != nil {
		return tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}

	testID = "#7JGP0R3o" // Invalid the Refresh Token because it was canceled in the #Gq7dHBlM test
	tests++
	_, _, _, err = test.refreshAccessToken(env,
		refreshToken,
		http.MethodPut,
		"/v1/oauth/token",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		404,
		false)
	if err != nil {
		return tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)

	return tests, nil

}

func (test *tests) revokeRefreshToken(env *c.Env, accessToken, publicSessionID, method, endpoint,
	headerContentType, headerRequestID string, statusCodeExpected int) error {

	var (
		err     error
		headers map[string]string
		props   struct {
			PublicSessionID string `json:"sessionID"`
		}
		statusCode int
	)

	headers = make(map[string]string)
	if len(headerContentType) > 0 {
		headers["content-type"] = headerContentType
	}
	if len(headerRequestID) > 0 {
		headers["request-id"] = headerRequestID
	}
	if len(accessToken) > 0 {
		headers["authorization"] = fmt.Sprintf("bearer %s", accessToken)
	}

	props.PublicSessionID = publicSessionID

	statusCode, err = u.HttpClient(env, method, endpoint, headers, props, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}
