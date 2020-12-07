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

func (checkList *checkLists) getListOfActiveSession(env *c.Env, accessToken, publicSessionID string) (int, error) {

	const checkListName = "Get a list of active sessions"
	const indentBeforeStatus = "\t\t\t\t\t\t"
	const errorMessage = "Failed to complete the %s test %s [%s]"
	var err error
	var tests = 0
	var testID string

	fmt.Print(g.ItemII)
	fmt.Print(checkListName, indentBeforeStatus)

	testID = "#MOi0O3GS" // No the Access Token
	err = test.getListOfActiveSessions(env,
		"",
		publicSessionID,
		http.MethodGet,
		"/v1/oauth/sessions",
		"K7800-H7625-Z5852-N1693-K1972",
		401,
		false)
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#jg7U4TDR" // Invalid the Access Token
	err = test.getListOfActiveSessions(env,
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJGaW5hbmNlbGltZS5jb20iLCJzdWIiOiJBdXRob3JpemF0aW9uIiwicHVycG9zZSI6ImFjY2VzcyIsImlkIjoiMTZkN2RiNTM3MjQ3ZWFmMTEzZjRjOGFkNTllOWEyYTU4OWNlN2NhZjYxMzViY2Q3YmZlYzBiNTI1YWY0OGEyYSIsImlhdCI6MTU5NjgzNTM5OX0.d68bea3232f10c60483f838fff8d8c66661cb497b141213c9a006be2e7c9d723",
		publicSessionID,
		http.MethodGet,
		"/v1/oauth/sessions",
		"K7800-H7625-Z5852-N1693-K1972",
		401,
		false)
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#VaK6YsxC" // No the RequestID header
	err = test.getListOfActiveSessions(env,
		accessToken,
		publicSessionID,
		http.MethodGet,
		"/v1/oauth/sessions",
		"",
		400,
		false)
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#6R5YXIzJ" // Invalid the POST method
	err = test.getListOfActiveSessions(env,
		accessToken,
		publicSessionID,
		http.MethodPost,
		"/v1/oauth/sessions",
		"K7800-H7625-Z5852-N1693-K1972",
		405,
		false)
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#KnGUxH36" // OK
	err = test.getListOfActiveSessions(env,
		accessToken,
		publicSessionID,
		http.MethodGet,
		"/v1/oauth/sessions",
		"K7800-H7625-Z5852-N1693-K1972",
		200,
		true)
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)

	return tests, nil

}

func (test *tests) getListOfActiveSessions(env *c.Env, accessToken, publicSessionID, method, endpoint,
	headerRequestID string, statusCodeExpected int, checkResponseValues bool) error {

	var (
		err        error
		headers    map[string]string
		response   []session
		statusCode int
	)

	headers = make(map[string]string)
	if len(headerRequestID) > 0 {
		headers["request-id"] = headerRequestID
	}
	if len(accessToken) > 0 {
		headers["authorization"] = fmt.Sprintf("bearer %s", accessToken)
	}

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, &response)
	if err != nil {
		return errors.New(fmt.Sprintf("Faled to get a list of active sessions [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	if checkResponseValues {

		if len(response) == 0 {
			return errors.New("REST-API service returned empty response")
		}

		if response[0].SessionID != publicSessionID {
			return errors.New(fmt.Sprintf("REST-API service returned the wrong sessionID value: got %v want %v",
				response[0].SessionID, publicSessionID))
		}
	}

	return nil
}
