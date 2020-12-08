package authorization

import (
	"errors"
	"fmt"
	c "github.com/dmalix/financelime-functional-tests/cmd/config"
	g "github.com/dmalix/financelime-functional-tests/cmd/pgraphics"
	u "github.com/dmalix/financelime-functional-tests/cmd/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (checkList *checkLists) confirmUserEmail(env *c.Env) (int, error) {

	const checkListName = "Confirm User Email (+30 sec of waiting)"
	const indentBeforeStatus = "\t\t\t\t"
	const errorMessage = "Failed to complete the %s test %s [%s]"
	var err error
	var tests = 0
	var testID string
	var confirmationKey string

	fmt.Print(g.ItemII)
	fmt.Print(checkListName, indentBeforeStatus)

	testID = "#fF1iSKo8" // OK
	confirmationKey, err = test.getConfirmationKeyFromEmailMessage(env,
		&u.IMAP{
			Host:     env.Config.IMAP.Host,
			Port:     env.Config.IMAP.Port,
			Username: env.Config.IMAP.Users.Username[1],
			Password: env.Config.IMAP.Users.Password[1],
		})
	if err != nil {
		return tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#G7zv5jx4" // Invalid ENDPOINT
	err = test.confirmUserEmailWithConfirmationKey(env,
		"",
		http.MethodGet,
		"/u/123456789",
		404)
	if err != nil {
		return tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#1iSKofF8" // OK
	err = test.confirmUserEmailWithConfirmationKey(env,
		confirmationKey,
		http.MethodGet,
		"/u/%s",
		200)
	if err != nil {
		return tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#eJ7RFkI5" // Not Found
	err = test.confirmUserEmailWithConfirmationKey(env,
		confirmationKey,
		http.MethodGet,
		"/u/%s",
		404)
	if err != nil {
		return tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	testID = "#DwStF6M0"
	if err = signUp_409_UserAlreadyExist2(env); err != nil {
		return tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}
	tests++

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)

	return tests, nil
}

func (test *tests) getConfirmationKeyFromEmailMessage(env *c.Env, imap *u.IMAP) (string, error) {

	var (
		messageIdList         []string
		messageID             string
		err                   error
		confirm               bool
		confirmationKeySource string
		confirmationKeyResult string
		domain                string
		messageIDSplit        []string
		temp                  []string
	)

	time.Sleep(30 * time.Second)

	messageIdList, err = imap.GetMailHeaders()
	if err != nil {
		return confirmationKeyResult, errors.New(fmt.Sprintf("Failed to get messageId list [%s]", err))
	}

	for i := 0; i <= len(messageIdList)-1; i++ {

		if confirm {
			continue
		}

		messageID = messageIdList[len(messageIdList)-1-i]
		messageIDSplit = strings.Split(messageID, "@")
		if len(messageIDSplit) != 2 {
			continue
		}

		// Get KeyLink
		temp = strings.Split(messageIDSplit[0], "<")
		if len(temp) != 2 {
			continue
		}
		confirmationKeySource = temp[1]

		// Get Domain
		temp = strings.Split(messageIDSplit[1], ">")
		if len(temp) != 2 {
			continue
		}
		domain = temp[0]
		if fmt.Sprintf("%s.%s", "sign-up", env.Config.General.Domain) != domain {
			continue
		}
		confirm = true
		confirmationKeyResult = confirmationKeySource
	}

	if confirm {
		return confirmationKeyResult, nil
	} else {
		return confirmationKeyResult,
			errors.New("failed to get a confirmationKey from the email")
	}
}

func (test *tests) confirmUserEmailWithConfirmationKey(env *c.Env, confirmationKey, method, endpoint string,
	statusCodeExpected int) error {

	var (
		err        error
		headers    map[string]string
		statusCode int
	)

	headers = make(map[string]string)

	if len(confirmationKey) > 0 {
		endpoint = fmt.Sprintf(endpoint, confirmationKey)
	}

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}
