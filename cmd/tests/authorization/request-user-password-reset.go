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

func (checkList *checkLists) requestUserPasswordReset(env *c.Env) (int, error) {

	const checkListName = "Request User Password reset"
	const indentBeforeStatus = "\t\t\t\t\t\t"
	const errorMessage = "Failed to complete the %s test %s [%s]"
	var err error
	var tests = 0
	var testID string

	fmt.Print(g.ItemII)
	fmt.Print(checkListName, indentBeforeStatus)

	testID = "#Afeich8a" // No the RequestID header
	err = test.requestUserPasswordReset(env,
		http.MethodPost,
		"/v1/resetpassword",
		"application/json;charset=utf-8",
		"",
		400,
		resetPassword{
			Email: env.Config.IMAP.Users.Username[1],
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#aec9yiB5" // No the ContentType header
	err = test.requestUserPasswordReset(env,
		http.MethodPost,
		"/v1/resetpassword",
		"",
		"K7800-H7625-Z5852-N1693-K1972",
		400,
		resetPassword{
			Email: env.Config.IMAP.Users.Username[1],
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#eeS0jae3" // Invalid User Email
	err = test.requestUserPasswordReset(env,
		http.MethodPost,
		"/v1/resetpassword",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		400,
		resetPassword{
			Email: "",
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#Da0og2Im" // User Not Found
	err = test.requestUserPasswordReset(env,
		http.MethodPost,
		"/v1/resetpassword",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		404,
		resetPassword{
			Email: "user@domain.com",
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#uboVe6sa" // OK
	err = test.requestUserPasswordReset(env,
		http.MethodPost,
		"/v1/resetpassword",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		204,
		resetPassword{
			Email: env.Config.IMAP.Users.Username[1],
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)

	return tests, nil
}

func (test *tests) requestUserPasswordReset(env *c.Env, method, endpoint,
	headerContentType, headerRequestID string, statusCodeExpected int, props resetPassword) error {

	var (
		err        error
		headers    map[string]string
		propsJSON  []byte
		statusCode int
	)

	headers = make(map[string]string)
	if len(headerContentType) > 0 {
		headers["content-type"] = headerContentType
	}
	if len(headerContentType) > 0 {
		headers["request-id"] = headerRequestID
	}

	statusCode, err = u.HttpClient(env, method, endpoint, headers, props, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		propsJSON, err = json.Marshal(props)
		if err != nil {
			return errors.New("failed to convert props data to JSON-format")
		}
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v [props: %s]",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected), string(propsJSON)))
	}

	return nil
}
