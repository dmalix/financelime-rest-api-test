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

/*	Return:
	sessionID    string
	accessToken	 string
	refreshToken string
	err          error
*/
func requestAccessToken_200(env *c.Env, pPassword string) (string, string, string, error) {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/oauth/token/request"
		statusCodeExpected = 200
	)

	var (
		err          error
		headers      map[string]string
		props        getAccessToken
		response     jwt
		responseJson []byte
		statusCode   int
		sessionID    string
		accessToken  string
		refreshToken string
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.Email = env.Config.IMAP.Users.Username[1]
	props.Password = pPassword
	props.Device.Language = "en-US"
	props.Device.Platform = "Linux x86_64"
	props.Device.Timezone = "2"
	props.Device.Height = 1920
	props.Device.Width = 1060

	statusCode, err = u.HttpClient(env, method, endpoint, headers, props, &response)
	if err != nil {
		return sessionID, accessToken, refreshToken,
			errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return sessionID, accessToken, refreshToken,
			errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
				strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	if len(response.SessionID) == 0 || len(response.AccessToken) == 0 || len(response.RefreshToken) == 0 {
		responseJson, err = json.Marshal(response)
		if err != nil {
			return sessionID, accessToken, refreshToken,
				errors.New("failed to convert response data to JSON-format")
		}
		return sessionID, accessToken, refreshToken,
			errors.New(fmt.Sprintf("REST-API service returned empty value: [%s, %s, %s] %s",
				strconv.Itoa(len(response.SessionID)),
				strconv.Itoa(len(response.AccessToken)),
				strconv.Itoa(len(response.RefreshToken)),
				string(responseJson)))
	}

	sessionID = response.SessionID
	accessToken = response.AccessToken
	refreshToken = response.RefreshToken

	return sessionID, accessToken, refreshToken, nil
}

func getEmailAboutLoginAction(env *c.Env) error {

	var (
		imap                    u.IMAP
		messageIdList           []string
		messageID               string
		err                     error
		confirm                 bool
		remoteAddressSourceList []string
		remoteAddressSource     string
		domainList              []string
		domain                  string
		messageIDSplit          []string
		temp                    []string
		domainExpected          = fmt.Sprintf("%s.%s", "request-access-token", env.Config.General.Domain)
	)

	time.Sleep(30 * time.Second)

	imap.Host = env.Config.IMAP.Host
	imap.Port = env.Config.IMAP.Port
	imap.Username = env.Config.IMAP.Users.Username[1]
	imap.Password = env.Config.IMAP.Users.Password[1]

	messageIdList, err = imap.GetMailHeaders()
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to get messageId list [%s]", err))
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
		remoteAddressSource = temp[1]

		// Get Domain
		temp = strings.Split(messageIDSplit[1], ">")
		if len(temp) != 2 {
			continue
		}
		domain = temp[0]

		domainList = append(domainList, domain)
		remoteAddressSourceList = append(remoteAddressSourceList, remoteAddressSource)

		if domainExpected != domain {
			continue
		}
		if remoteAddressSource != env.Config.CurrentRemoteAddress {
			continue
		}

		confirm = true
	}

	if confirm {
		return nil
	} else {
		return errors.New(
			fmt.Sprintf(
				"Not found an email about a login action (got %s want %s) or an email has wrong the current remote address (got %s want %s)",
				strings.Join(domainList, ","), domainExpected,
				strings.Join(remoteAddressSourceList, ","), env.Config.CurrentRemoteAddress))
	}
}

func getListActiveSessions_401_NoHeaderAuthorization(env *c.Env) error {

	const (
		method             = http.MethodGet
		endpoint           = "/v1/oauth/sessions"
		statusCodeExpected = 401
	)

	var (
		err        error
		headers    map[string]string
		statusCode int
	)

	headers = make(map[string]string)
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("Faled to get a list of active sessions [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func getListActiveSessions_401_InvalidAccessToken(env *c.Env, pAccessToken string) error {

	const (
		method             = http.MethodGet
		endpoint           = "/v1/oauth/sessions"
		statusCodeExpected = 401
	)

	var (
		err        error
		headers    map[string]string
		statusCode int
	)

	headers = make(map[string]string)
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"
	headers["authorization"] = fmt.Sprintf(
		"bearer %s", pAccessToken)

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("Faled to get a list of active sessions [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func getListActiveSessions_401_NoBearerValueIntoHeaderAuthorization(env *c.Env, pAccessToken string) error {

	const (
		method             = http.MethodGet
		endpoint           = "/v1/oauth/sessions"
		statusCodeExpected = 401
	)

	var (
		err        error
		headers    map[string]string
		response   []session
		statusCode int
	)

	headers = make(map[string]string)
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"
	headers["authorization"] = fmt.Sprintf("%s", pAccessToken)

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, &response)
	if err != nil {
		return errors.New(fmt.Sprintf("Faled to get a list of active sessions [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func getListActiveSessions_400_NoHeaderRequestID(env *c.Env, pAccessToken string) error {

	const (
		method             = http.MethodGet
		endpoint           = "/v1/oauth/sessions"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		response   []session
		statusCode int
	)

	headers = make(map[string]string)
	headers["authorization"] = fmt.Sprintf("bearer %s", pAccessToken)

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, &response)
	if err != nil {
		return errors.New(fmt.Sprintf("Faled to get a list of active sessions [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func getListActiveSessions_405_InvalidMethodPost(env *c.Env, pAccessToken string) error {

	const (
		method             = http.MethodPost
		endpoint           = "/v1/oauth/sessions"
		statusCodeExpected = 405
	)

	var (
		err        error
		headers    map[string]string
		response   []session
		statusCode int
	)

	headers = make(map[string]string)
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"
	headers["authorization"] = fmt.Sprintf("bearer %s", pAccessToken)

	statusCode, err = u.HttpClient(env, method, endpoint, headers, nil, &response)
	if err != nil {
		return errors.New(fmt.Sprintf("Faled to get a list of active sessions [%s]", err))
	}
	if statusCode != statusCodeExpected {
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

func refreshAccessToken_403_InvalidRefreshToken(env *c.Env) error {

	const (
		method             = http.MethodPut
		endpoint           = "/account/oauth/token"
		statusCodeExpected = 403
	)

	var (
		err        error
		headers    map[string]string
		props      refreshToken
		response   jwt
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.RefreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJGaW5hbmNlbGltZS5jb20iLCJzdWIiOiJBdXRob3JpemF0aW9uIiwicHVycG9zZSI6InJlZnJlc2giLCJpZCI6IjE2ZDdkYjUzNzI0N2VhZjExM2Y0YzhhZDU5ZTlhMmE1ODljZTdjYWY2MTM1YmNkN2JmZWMwYjUyNWFmNDhhMmEiLCJpYXQiOjE1OTY4MzUzOTl9.b88345d361482865f1a7af533d41d66e922dcca4c76c2d4b1fcfa65616679471"

	statusCode, err = u.HttpClient(env, method, endpoint, headers, props, &response)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func refreshAccessToken_400_NoHeaderContentType(env *c.Env, pRefreshToken string) error {

	const (
		method             = http.MethodPut
		endpoint           = "/account/oauth/token"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		props      refreshToken
		response   jwt
		statusCode int
	)

	headers = make(map[string]string)
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.RefreshToken = pRefreshToken

	statusCode, err = u.HttpClient(env, method, endpoint, headers, props, &response)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func refreshAccessToken_400_NoHeaderRequestID(env *c.Env, pRefreshToken string) error {

	const (
		method             = http.MethodPut
		endpoint           = "/account/oauth/token"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		props      refreshToken
		response   jwt
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"

	props.RefreshToken = pRefreshToken

	statusCode, err = u.HttpClient(env, method, endpoint, headers, props, &response)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func refreshAccessToken_400_InvalidMethodPost(env *c.Env, pRefreshToken string) error {

	const (
		method             = http.MethodPost
		endpoint           = "/account/oauth/token"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		props      refreshToken
		response   jwt
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.RefreshToken = pRefreshToken

	statusCode, err = u.HttpClient(env, method, endpoint, headers, props, &response)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

func refreshAccessToken_405_InvalidMethodGet(env *c.Env, pRefreshToken string) error {

	const (
		method             = http.MethodGet
		endpoint           = "/account/oauth/token"
		statusCodeExpected = 405
	)

	var (
		err        error
		headers    map[string]string
		props      refreshToken
		response   jwt
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.RefreshToken = pRefreshToken

	statusCode, err = u.HttpClient(env, method, endpoint, headers, props, &response)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to complete the request [%s]", err))
	}
	if statusCode != statusCodeExpected {
		return errors.New(fmt.Sprintf("REST-API service returned wrong status code: got %v want %v",
			strconv.Itoa(statusCode), strconv.Itoa(statusCodeExpected)))
	}

	return nil
}

//	Return:
//		publicSessionID string
//		accessToken string
//		refreshToken string
//		error
func refreshAccessToken_200(env *c.Env, pRefreshToken string) (string, string, string, error) {

	const (
		method             = http.MethodPut
		endpoint           = "/account/oauth/token"
		statusCodeExpected = 200
	)

	var (
		err             error
		headers         map[string]string
		props           refreshToken
		response        jwt
		responseJson    []byte
		statusCode      int
		publicSessionID string
		accessToken     string
		refreshToken    string
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"

	props.RefreshToken = pRefreshToken

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

	if len(response.AccessToken) == 0 || len(response.RefreshToken) == 0 {
		responseJson, err = json.Marshal(response)
		if err != nil {
			return publicSessionID, accessToken, refreshToken,
				errors.New("failed to convert response data to JSON-format")
		}
		return publicSessionID, accessToken, refreshToken,
			errors.New(fmt.Sprintf("REST-API service returned empty value [%v, %v]: %s",
				strconv.Itoa(len(response.AccessToken)),
				strconv.Itoa(len(response.RefreshToken)),
				string(responseJson)))
	}

	accessToken = response.AccessToken
	refreshToken = response.RefreshToken

	return publicSessionID, accessToken, refreshToken, nil
}

func revokeAccessToken_400_NoHeaderRequestID(env *c.Env, pAccessToken string, pPublicSessionID string) error {

	const (
		method             = http.MethodPost
		endpoint           = "/account/oauth/revoke"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		props      revokeAccessToken
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["authorization"] = fmt.Sprintf("bearer %s", pAccessToken)

	props.PublicSessionID = pPublicSessionID

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

func revokeAccessToken_400_NoHeaderContentType(env *c.Env, pAccessToken string, pPublicSessionID string) error {

	const (
		method             = http.MethodPost
		endpoint           = "/account/oauth/revoke"
		statusCodeExpected = 400
	)

	var (
		err        error
		headers    map[string]string
		props      revokeAccessToken
		statusCode int
	)

	headers = make(map[string]string)
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"
	headers["authorization"] = fmt.Sprintf("bearer %s", pAccessToken)

	props.PublicSessionID = pPublicSessionID

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

func revokeAccessToken_405_InvalidMethodPut(env *c.Env, pAccessToken string, pPublicSessionID string) error {

	const (
		method             = http.MethodPut
		endpoint           = "/account/oauth/revoke"
		statusCodeExpected = 405
	)

	var (
		err        error
		headers    map[string]string
		props      revokeAccessToken
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"
	headers["authorization"] = fmt.Sprintf("bearer %s", pAccessToken)

	props.PublicSessionID = pPublicSessionID

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

func revokeAccessToken_405_InvalidMethodGet(env *c.Env, pAccessToken string, pPublicSessionID string) error {

	const (
		method             = http.MethodGet
		endpoint           = "/account/oauth/revoke"
		statusCodeExpected = 405
	)

	var (
		err        error
		headers    map[string]string
		props      revokeAccessToken
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"
	headers["authorization"] = fmt.Sprintf("bearer %s", pAccessToken)

	props.PublicSessionID = pPublicSessionID

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

func revokeAccessToken_200(env *c.Env, pAccessToken string, pPublicSessionID string) error {

	const (
		method             = http.MethodPost
		endpoint           = "/account/oauth/revoke"
		statusCodeExpected = 200
	)

	var (
		err        error
		headers    map[string]string
		props      revokeAccessToken
		statusCode int
	)

	headers = make(map[string]string)
	headers["content-type"] = "application/json;charset=utf-8"
	headers["request-id"] = "K7800-H7625-Z5852-N1693-K1972"
	headers["authorization"] = fmt.Sprintf("bearer %s", pAccessToken)

	props.PublicSessionID = pPublicSessionID

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
