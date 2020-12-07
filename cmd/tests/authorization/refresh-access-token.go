package authorization

import (
	"encoding/json"
	"errors"
	"fmt"
	c "github.com/dmalix/financelime-functional-tests/cmd/config"
	g "github.com/dmalix/financelime-functional-tests/cmd/pgraphics"
	u "github.com/dmalix/financelime-functional-tests/cmd/utils"
	"net/http"
	"strconv"
)

func (checkList *checkLists) refreshAccessToken(env *c.Env, refreshToken,
	publicSessionID string) (string, string, int, error) {

	const checkListName = "Refresh Access Token"
	const indentBeforeStatus = "\t\t\t\t\t\t\t"
	const errorMessage = "Failed to complete the %s test %s [%s]"
	var err error
	var tests = 0
	var testID string

	var accessToken string

	fmt.Print(g.ItemII)
	fmt.Print(checkListName, indentBeforeStatus)

	testID = "#pWs8HJ4F" // Invalid Refresh Token
	_, _, _, err = test.refreshAccessToken(env,
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJGaW5hbmNlbGltZS5jb20iLCJzdWIiOiJBdXRob3JpemF0aW9uIiwicHVycG9zZSI6InJlZnJlc2giLCJpZCI6IjE2ZDdkYjUzNzI0N2VhZjExM2Y0YzhhZDU5ZTlhMmE1ODljZTdjYWY2MTM1YmNkN2JmZWMwYjUyNWFmNDhhMmEiLCJpYXQiOjE1OTY4MzUzOTl9.b88345d361482865f1a7af533d41d66e922dcca4c76c2d4b1fcfa65616679471",
		http.MethodPut,
		"/v1/oauth/token",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		404,
		false)
	if err != nil {
		return accessToken, refreshToken, tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#Xhca9hTB" // No the ContentType header
	_, _, _, err = test.refreshAccessToken(env,
		"",
		http.MethodPut,
		"/v1/oauth/token",
		"",
		"K7800-H7625-Z5852-N1693-K1972",
		400,
		false)
	if err != nil {
		return accessToken, refreshToken, tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#cdWR7DgQ" // No the RequestID Header
	_, _, _, err = test.refreshAccessToken(env,
		"",
		http.MethodPut,
		"/v1/oauth/token",
		"application/json;charset=utf-8",
		"",
		400,
		false)
	if err != nil {
		return accessToken, refreshToken, tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#EDOv5oJ6" // Invalid the GET method
	_, _, _, err = test.refreshAccessToken(env,
		"",
		http.MethodGet,
		"/v1/oauth/token",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		405,
		false)
	if err != nil {
		return accessToken, refreshToken, tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#Wm1PBdVN" // Invalid the DELETE method
	_, _, _, err = test.refreshAccessToken(env,
		"",
		http.MethodDelete,
		"/v1/oauth/token",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		405,
		false)
	if err != nil {
		return accessToken, refreshToken, tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#9CGhxsWz" // OK
	_, accessToken, refreshToken, err = test.refreshAccessToken(env,
		refreshToken,
		http.MethodPut,
		"/v1/oauth/token",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		200,
		true)
	if err != nil {
		return accessToken, refreshToken, tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#n4XHQkCf"
	if err = getListActiveSessions_200(env, accessToken, publicSessionID); err != nil {
		return accessToken, refreshToken, tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)

	return accessToken, refreshToken, tests, nil
}

func (test *tests) refreshAccessToken(env *c.Env, refreshToken, method, endpoint,
	headerContentType, headerRequestID string, statusCodeExpected int, checkResponseValues bool) (string, string, string, error) {

	var (
		err     error
		headers map[string]string
		props   struct {
			RefreshToken string `json:"refreshToken"`
		}
		response        jwt
		responseJson    []byte
		statusCode      int
		publicSessionID string
		jwtAccess       string
		jwtRefresh      string
	)

	headers = make(map[string]string)
	if len(headerContentType) > 0 {
		headers["content-type"] = headerContentType
	}
	if len(headerRequestID) > 0 {
		headers["request-id"] = headerRequestID
	}

	props.RefreshToken = refreshToken

	statusCode, err = u.HttpClient(env, method, endpoint, headers, props, &response)
	if err != nil {
		return publicSessionID, jwtAccess, jwtRefresh,
			errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return publicSessionID, jwtAccess, jwtRefresh,
			errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
				strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	if checkResponseValues {

		if len(response.PublicSessionID) == 0 || len(response.AccessToken) == 0 || len(response.RefreshToken) == 0 {
			responseJson, err = json.Marshal(response)
			if err != nil {
				return publicSessionID, jwtAccess, jwtRefresh,
					errors.New("failed to convert response data to JSON-format")
			}
			return publicSessionID, jwtAccess, jwtRefresh,
				errors.New(fmt.Sprintf("REST-API service returned empty value [%d, %d, %d]: %s",
					len(response.PublicSessionID),
					len(response.AccessToken),
					len(response.RefreshToken),
					string(responseJson)))
		}

	}

	jwtAccess = response.AccessToken
	jwtRefresh = response.RefreshToken
	publicSessionID = response.PublicSessionID

	return publicSessionID, jwtAccess, jwtRefresh, nil
}
