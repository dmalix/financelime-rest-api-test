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

func (checkList *checkLists) signUp(env *c.Env) (int, error) {

	const checkListName = "Sign Up"
	const indentBeforeStatus = "\t\t\t\t\t\t\t\t"
	const errorMessage = "Failed to complete the %s test %s [%s]"
	var err error
	var tests = 0
	var testID string

	fmt.Print(g.ItemII)
	fmt.Print(checkListName, indentBeforeStatus)

	testID = "#hdpFBOG7" // No the RequestID header
	err = test.signUp(env,
		http.MethodPost,
		"/v1/signup",
		"application/json;charset=utf-8",
		"",
		400,
		signup{
			Email:      env.Config.IMAP.Users.Username[1],
			Language:   "ru",
			InviteCode: env.Config.InviteCode,
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#VuLq72Fi" // No the ContentType header
	err = test.signUp(env,
		http.MethodPost,
		"/v1/signup",
		"",
		"K7800-H7625-Z5852-N1693-K1972",
		400,
		signup{
			Email:      env.Config.IMAP.Users.Username[1],
			Language:   "ru",
			InviteCode: env.Config.InviteCode,
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#T6PwbuNa" // Invalid Language
	err = test.signUp(env,
		http.MethodPost,
		"/v1/signup",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		400,
		signup{
			Email:      env.Config.IMAP.Users.Username[1],
			Language:   "de",
			InviteCode: env.Config.InviteCode,
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#02GXjc2C" // Invalid Language
	err = test.signUp(env,
		http.MethodPost,
		"/v1/signup",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		400,
		signup{
			Email:      env.Config.IMAP.Users.Username[1],
			Language:   "1234",
			InviteCode: env.Config.InviteCode,
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#yBo3E2X4" // Invalid Language
	err = test.signUp(env,
		http.MethodPost,
		"/v1/signup",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		400,
		signup{
			Email:      env.Config.IMAP.Users.Username[1],
			Language:   "",
			InviteCode: env.Config.InviteCode,
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#9C73chj1" // Invalid User Email
	err = test.signUp(env,
		http.MethodPost,
		"/v1/signup",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		400,
		signup{
			Email:      "",
			Language:   "ru",
			InviteCode: env.Config.InviteCode,
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#tc9Id2BU" // Invalid Invite Code
	err = test.signUp(env,
		http.MethodPost,
		"/v1/signup",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		409,
		signup{
			Email:      env.Config.IMAP.Users.Username[1],
			Language:   "ru",
			InviteCode: "1234567890",
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#VDR1Hmzv" // User Already Exist (0)
	err = test.signUp(env,
		http.MethodPost,
		"/v1/signup",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		409,
		signup{
			Email:      env.Config.IMAP.Users.Username[0],
			Language:   "ru",
			InviteCode: env.Config.InviteCode,
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#nRIxV62A" // OK (1)
	err = test.signUp(env,
		http.MethodPost,
		"/v1/signup",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		204,
		signup{
			Email:      env.Config.IMAP.Users.Username[1],
			Language:   "ru",
			InviteCode: env.Config.InviteCode,
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#dREV5jQx" // User Already Exist (1)
	err = test.signUp(env,
		http.MethodPost,
		"/v1/signup",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		409,
		signup{
			Email:      env.Config.IMAP.Users.Username[1],
			Language:   "ru",
			InviteCode: env.Config.InviteCode,
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#Dyr5SSNC" // OK (2)
	err = test.signUp(env,
		http.MethodPost,
		"/v1/signup",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		204,
		signup{
			Email:      env.Config.IMAP.Users.Username[2],
			Language:   "en",
			InviteCode: env.Config.InviteCode,
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#DwStF6M0" // User Already Exist (2)
	err = test.signUp(env,
		http.MethodPost,
		"/v1/signup",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		409,
		signup{
			Email:      env.Config.IMAP.Users.Username[2],
			Language:   "en",
			InviteCode: env.Config.InviteCode,
		})
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)

	return tests, nil
}

func (test *tests) signUp(env *c.Env, method, endpoint,
	headerContentType, headerRequestID string, statusCodeExpected int, props signup) error {

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
