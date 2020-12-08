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

func (checkList *checkLists) requestAccessToken(env *c.Env, accessToken, refreshToken, publicSessionID,
	newPassword string) (string, string, string, int, error) {

	const checkListName = "Request Access Token"
	const indentBeforeStatus = "\t\t\t\t\t\t\t"
	const errorMessage = "Failed to complete the %s test %s [%s]"
	var err error
	var tests = 0
	var testID string

	fmt.Print(g.ItemII)
	fmt.Print(checkListName, indentBeforeStatus)

	testID = "#Kp0gJ1mm"
	accessToken, refreshToken, publicSessionID, err = test.requestAccessToken(env,
		http.MethodPost,
		"/v1/oauth/token",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		200,
		&requestAccessToken{
			Email:    env.Config.IMAP.Users.Username[1],
			Password: newPassword,
			Device: device{
				Language: "en-US",
				Platform: "Linux x86_64",
				Timezone: "2",
				Height:   1920,
				Width:    1060,
			},
		},
		true)
	if err != nil {
		return accessToken, refreshToken, publicSessionID, tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)

	return accessToken, refreshToken, publicSessionID, tests, nil
}

func (test *tests) requestAccessToken(env *c.Env, method, endpoint, headerContentType,
	headerRequestID string, statusCodeExpected int, props *requestAccessToken, checkResponseValues bool) (string, string, string, error) {

	var (
		err             error
		headers         map[string]string
		response        jwt
		responseJson    []byte
		statusCode      int
		publicSessionID string
		accessToken     string
		refreshToken    string
	)

	headers = make(map[string]string)
	headers["content-type"] = headerContentType
	headers["request-id"] = headerRequestID

	statusCode, err = u.HttpClient(env, method, endpoint, headers, props, &response)
	if err != nil {
		return publicSessionID, accessToken, refreshToken,
			errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return publicSessionID, accessToken, refreshToken,
			errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
				strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	if checkResponseValues {

		if len(response.PublicSessionID) == 0 || len(response.AccessToken) == 0 || len(response.RefreshToken) == 0 {
			responseJson, err = json.Marshal(response)
			if err != nil {
				return publicSessionID, accessToken, refreshToken,
					errors.New("failed to convert response data to JSON-format")
			}
			return publicSessionID, accessToken, refreshToken,
				errors.New(fmt.Sprintf("REST-API service returned empty value: [%s, %s, %s] %s",
					strconv.Itoa(len(response.PublicSessionID)),
					strconv.Itoa(len(response.AccessToken)),
					strconv.Itoa(len(response.RefreshToken)),
					string(responseJson)))
		}

		publicSessionID = response.PublicSessionID
		accessToken = response.AccessToken
		refreshToken = response.RefreshToken
	}

	return accessToken, refreshToken, publicSessionID, nil
}
