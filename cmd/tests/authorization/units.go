package authorization

import (
	"encoding/json"
	"errors"
	"fmt"
	c "github.com/dmalix/financelime-functional-tests/cmd/config"
	u "github.com/dmalix/financelime-functional-tests/cmd/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func signUp_400_NoHeaderRequestID(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/signup"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		props      signup
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"

	props.Email = env.Config.IMAP.Users.Username[0]
	props.Language = "en"
	props.InviteCode = env.Config.InviteCode

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

func signUp_400_NoHeaderContentType(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/signup"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		props      signup
		statusCode int
	)

	headers = make(map[string]string)
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = env.Config.IMAP.Users.Username[0]
	props.Language = "en"
	props.InviteCode = env.Config.InviteCode

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

func signUp_400_InvalidLanguageDe(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/signup"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		props      signup
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = env.Config.IMAP.Users.Username[0]
	props.Language = "de"
	props.InviteCode = env.Config.InviteCode

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

func signUp_400_InvalidLanguage1234(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/signup"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		props      signup
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = env.Config.IMAP.Users.Username[0]
	props.Language = "1234"
	props.InviteCode = env.Config.InviteCode

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

func signUp_400_InvalidLanguage(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/signup"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		props      signup
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = env.Config.IMAP.Users.Username[0]
	props.Language = ""
	props.InviteCode = env.Config.InviteCode

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

func signUp_400_InvalidEmail(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/signup"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		props      signup
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = ""
	props.Language = "ru"
	props.InviteCode = env.Config.InviteCode

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

func signUp_409_InvalidInviteCode(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/signup"
		statusCodeExpected = 409
	)

	var (
		err        error
		headers    map[string]string
		props      signup
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = env.Config.IMAP.Users.Username[0]
	props.Language = "ru"
	props.InviteCode = "1234567890"

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

func signUp_409_UserAlreadyExist0(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/signup"
		statusCodeExpected = 409
	)

	var (
		err        error
		headers    map[string]string
		props      signup
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = env.Config.IMAP.Users.Username[0]
	props.Language = "ru"
	props.InviteCode = env.Config.InviteCode

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

func signUp_204_Email1(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/signup"
		statusCodeExpected = 204
	)

	var (
		err        error
		headers    map[string]string
		props      signup
		propsJSON  []byte
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = env.Config.IMAP.Users.Username[1]
	props.Language = "ru"
	props.InviteCode = env.Config.InviteCode

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

func signUp_409_UserAlreadyExist1(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/signup"
		statusCodeExpected = 409
	)

	var (
		err        error
		headers    map[string]string
		props      signup
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = env.Config.IMAP.Users.Username[1]
	props.Language = "ru"
	props.InviteCode = env.Config.InviteCode

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

func signUp_204_Email2(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/signup"
		statusCodeExpected = 204
	)

	var (
		err        error
		headers    map[string]string
		props      signup
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = env.Config.IMAP.Users.Username[2]
	props.Language = "en"
	props.InviteCode = env.Config.InviteCode

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

func signUp_409_UserAlreadyExist2(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/signup"
		statusCodeExpected = 409
	)

	var (
		err        error
		headers    map[string]string
		props      signup
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = env.Config.IMAP.Users.Username[2]
	props.Language = "en"
	props.InviteCode = env.Config.InviteCode

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

func confirmEmail_404_InvalidEndPoint123456789(env *c.Env) error {

	const (
		method             = http.MethodGet
		endpoint           = "/u/123456789"
		statusCodeExpected = 404
	)

	var (
		err        error
		headers    map[string]string
		statusCode int
	)

	headers = make(map[string]string)

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

func confirmEmail_200(env *c.Env) (string, error) {

	const (
		method             = http.MethodGet
		endpoint           = "/u/%s"
		statusCodeExpected = 200
	)

	var (
		imap                  u.IMAP
		messageIdList         []string
		messageID             string
		err                   error
		statusCodeList        []string
		confirm               bool
		statusCode            int
		confirmationKeySource string
		confirmationKeyResult string
		domain                string
		messageIDSplit        []string
		temp                  []string
	)

	time.Sleep(30 * time.Second)

	imap.Host = env.Config.IMAP.Host
	imap.Port = env.Config.IMAP.Port
	imap.Username = env.Config.IMAP.Users.Username[1]
	imap.Password = env.Config.IMAP.Users.Password[1]

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

		statusCode, err =
			u.HttpClient(env, method, fmt.Sprintf(endpoint, confirmationKeySource), nil, nil, nil)
		if err != nil {
			return confirmationKeyResult, errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
		}
		if statusCode == statusCodeExpected {
			confirm = true
			confirmationKeyResult = confirmationKeySource
		}
		statusCodeList = append(statusCodeList, strconv.Itoa(statusCode))
	}

	if confirm {
		return confirmationKeyResult, nil
	} else {
		return confirmationKeyResult,
			errors.New(fmt.Sprintf("Failed to confirm email [%s]", strings.Join(statusCodeList, ",")))
	}
}

func confirmEmail_404(env *c.Env, pLinkKey string) error {

	const (
		method             = http.MethodGet
		endpoint           = "/u/%s"
		statusCodeExpected = 404
	)

	var (
		err        error
		statusCode int
	)

	statusCode, err =
		u.HttpClient(env, method, fmt.Sprintf(endpoint, pLinkKey), nil, nil, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func getPasswordFromEmail(env *c.Env) (string, error) {

	var (
		imap           u.IMAP
		messageIdList  []string
		messageID      string
		err            error
		confirm        bool
		passwordSource string
		passwordResult string
		domain         string
		messageIDSplit []string
		temp           []string
	)

	time.Sleep(30 * time.Second)

	imap.Host = env.Config.IMAP.Host
	imap.Port = env.Config.IMAP.Port
	imap.Username = env.Config.IMAP.Users.Username[1]
	imap.Password = env.Config.IMAP.Users.Password[1]

	messageIdList, err = imap.GetMailHeaders()
	if err != nil {
		return passwordResult, errors.New(fmt.Sprintf("Failed to get messageId list [%s]", err))
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
		passwordSource = temp[1]

		// Get Domain
		temp = strings.Split(messageIDSplit[1], ">")
		if len(temp) != 2 {
			continue
		}
		domain = temp[0]
		if fmt.Sprintf("%s.%s", "confirm-user-email", env.Config.General.Domain) != domain {
			continue
		}
		confirm = true
		passwordResult = passwordSource
	}

	if confirm {
		return passwordResult, nil
	} else {
		return passwordResult,
			errors.New("failed to get a password from the email")
	}
}

func resetUserPassword_400_NoHeaderContentType(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/account/resetpassword"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		props      resetPassword
		statusCode int
	)

	headers = make(map[string]string)
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = env.Config.IMAP.Users.Username[1]

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

func resetUserPassword_400_NoHeaderRequestID(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/account/resetpassword"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		props      resetPassword
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"

	props.Email = env.Config.IMAP.Users.Username[2]

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

func resetUserPassword_202(env *c.Env) error {

	const (
		method             = http.MethodPost
		endpoint           = "/account/resetpassword"
		statusCodeExpected = 202
	)

	var (
		err        error
		headers    map[string]string
		props      resetPassword
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = env.Config.IMAP.Users.Username[1]

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

func confirmPasswordReset_200(env *c.Env) (string, error) {

	const (
		method             = http.MethodGet
		endpoint           = "/p/%s"
		statusCodeExpected = 200
	)

	var (
		imap           u.IMAP
		messageIdList  []string
		messageID      string
		err            error
		statusCodeList []string
		confirm        bool
		statusCode     int
		linkKeySource  string
		linkKeyResult  string
		domain         string
		messageIDSplit []string
		temp           []string
	)

	time.Sleep(30 * time.Second)

	imap.Host = env.Config.IMAP.Host
	imap.Port = env.Config.IMAP.Port
	imap.Username = env.Config.IMAP.Users.Username[1]
	imap.Password = env.Config.IMAP.Users.Password[1]

	messageIdList, err = imap.GetMailHeaders()
	if err != nil {
		return linkKeyResult, errors.New(fmt.Sprintf("Failed to get messageId list [%s]", err))
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
		linkKeySource = temp[1]

		// Get Domain
		temp = strings.Split(messageIDSplit[1], ">")
		if len(temp) != 2 {
			continue
		}
		domain = temp[0]
		if fmt.Sprintf("%s.%s", "password-reset", env.Config.General.Domain) != domain {
			continue
		}

		statusCode, err =
			u.HttpClient(env, method, fmt.Sprintf(endpoint, linkKeySource), nil, nil, nil)
		if err != nil {
			return linkKeyResult, errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
		}
		if statusCode == statusCodeExpected {
			confirm = true
			linkKeyResult = linkKeySource
		}
		statusCodeList = append(statusCodeList, strconv.Itoa(statusCode))
	}

	if confirm {
		return linkKeyResult, nil
	} else {
		return linkKeyResult,
			errors.New(fmt.Sprintf("Failed to Confirm password reset [%s]", strings.Join(statusCodeList, ",")))
	}
}

func confirmPasswordReset_404(env *c.Env, pLinkKey string) error {

	const (
		method             = http.MethodGet
		endpoint           = "/p/%s"
		statusCodeExpected = 404
	)

	var (
		err        error
		statusCode int
	)

	statusCode, err =
		u.HttpClient(env, method, fmt.Sprintf(endpoint, pLinkKey), nil, nil, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func getListActiveSessions_200(env *c.Env, accessToken, sessionID string) error {

	const (
		method             = http.MethodGet
		endpoint           = "/v1/oauth/sessions"
		statusCodeExpected = 200
	)

	var (
		err        error
		headers    map[string]string
		response   []session
		statusCode int
	)

	headers = make(map[string]string)
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"
	headers["authorization"] = fmt.Sprintf("bearer %s", accessToken)

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, &response)
	if err != nil {
		return errors.New(fmt.Sprintf("Faled to get a list of active sessions [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	if len(response) == 0 {
		return errors.New("REST-API service returned empty response")
	}

	if response[0].SessionID != sessionID {
		return errors.New(fmt.Sprintf("REST-API service returned the wrong sessionID value: got %v want %v",
			response[0].SessionID, sessionID))
	}

	return nil
}
