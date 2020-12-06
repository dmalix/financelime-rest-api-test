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

type checkLists int
type tests int
var checklist checkLists
var test tests

func RunTests(env *c.Env) (int, error) {

	var (
		err                error
		testName           string
		tests              int
		totalTests         int
		testID             string
		indentBeforeStatus string
		linkKey            string
		newPassword        string
		publicSessionID    string
		accessToken        string
		refreshToken       string
	)

	const errorMessage = "Failed to complete the test %s [%s]"

	fmt.Print(g.PointI0)
	u.Colorize(u.ColorYellow, "Authorization tests group", true)
	fmt.Println(g.SpaceII)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Sign Up"
	indentBeforeStatus = "\t\t\t\t\t\t\t\t"
	tests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#hdpFBOG7"
	tests++
	if err = signUp_400_NoHeaderRequestID(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#VuLq72Fi"
	tests++
	if err = signUp_400_NoHeaderContentType(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#T6PwbuNa"
	tests++
	if err = signUp_400_InvalidLanguageDe(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#02GXjc2C"
	tests++
	if err = signUp_400_InvalidLanguage1234(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#yBo3E2X4"
	tests++
	if err = signUp_400_InvalidLanguage(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#9C73chj1"
	tests++
	if err = signUp_400_InvalidEmail(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#tc9Id2BU"
	tests++
	if err = signUp_409_InvalidInviteCode(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#VDR1Hmzv"
	tests++
	if err = signUp_409_UserAlreadyExist0(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#nRIxV62A"
	tests++
	if err = signUp_204_Email1(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#dREV5jQx"
	tests++
	if err = signUp_409_UserAlreadyExist1(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#dREV5jQx"
	tests++
	if err = signUp_409_UserAlreadyExist1(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#Dyr5SSNC"
	tests++
	if err = signUp_204_Email2(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#DwStF6M0"
	tests++
	if err = signUp_409_UserAlreadyExist2(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)
	totalTests = totalTests + tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Confirm user email (+30 sec of waiting)"
	indentBeforeStatus = "\t\t\t\t"
	tests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#G7zv5jx4"
	tests++
	if err = confirmEmail_404_InvalidEndPoint123456789(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#1iSKofF8"
	tests++
	if linkKey, err = confirmEmail_200(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#eJ7RFkI5"
	tests++
	if err = confirmEmail_404(env, linkKey); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#DwStF6M0"
	tests++
	if err = signUp_409_UserAlreadyExist2(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)
	totalTests = totalTests + tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Get a password from the email (+30 sec of waiting)"
	indentBeforeStatus = "\t\t\t"
	tests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#F2ArtAXB"
	if newPassword, err = getPasswordFromEmail(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, "Done", true)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Request Access Token"
	indentBeforeStatus = "\t\t\t\t\t\t\t"
	tests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#Kp0gJ1mm"
	tests++
	if publicSessionID, accessToken, refreshToken, err = requestAccessToken_200(env, newPassword); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)
	totalTests = totalTests + tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Get an email about a login action (+30 sec of waiting)"
	indentBeforeStatus = "\t\t"
	tests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#piQaaBr0"
	if err = getEmailAboutLoginAction(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, "Done", true)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Get a list of active sessions"
	indentBeforeStatus = "\t\t\t\t\t\t"
	tests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#MOi0O3GS"
	tests++
	if err = getListActiveSessions_401_NoHeaderAuthorization(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#jg7U4TDR"
	tests++
	if err = getListActiveSessions_401_InvalidAccessToken(env, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJGaW5hbmNlbGltZS5jb20iLCJzdWIiOiJBdXRob3JpemF0aW9uIiwicHVycG9zZSI6ImFjY2VzcyIsImlkIjoiMTZkN2RiNTM3MjQ3ZWFmMTEzZjRjOGFkNTllOWEyYTU4OWNlN2NhZjYxMzViY2Q3YmZlYzBiNTI1YWY0OGEyYSIsImlhdCI6MTU5NjgzNTM5OX0.d68bea3232f10c60483f838fff8d8c66661cb497b141213c9a006be2e7c9d723"); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#C9a10Hr1"
	tests++
	if err = getListActiveSessions_401_NoBearerValueIntoHeaderAuthorization(env, accessToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#VaK6YsxC"
	tests++
	if err = getListActiveSessions_400_NoHeaderRequestID(env, accessToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#6R5YXIzJ"
	tests++
	if err = getListActiveSessions_405_InvalidMethodPost(env, accessToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#KnGUxH36"
	tests++
	if err = getListActiveSessions_200(env, accessToken, publicSessionID); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)
	totalTests = totalTests + tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Refresh Access Token"
	indentBeforeStatus = "\t\t\t\t\t\t\t"
	tests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#pWs8HJ4F" // Invalid Refresh Token
	tests++
	_, _, _, err = refreshAccessToken(env,
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJGaW5hbmNlbGltZS5jb20iLCJzdWIiOiJBdXRob3JpemF0aW9uIiwicHVycG9zZSI6InJlZnJlc2giLCJpZCI6IjE2ZDdkYjUzNzI0N2VhZjExM2Y0YzhhZDU5ZTlhMmE1ODljZTdjYWY2MTM1YmNkN2JmZWMwYjUyNWFmNDhhMmEiLCJpYXQiOjE1OTY4MzUzOTl9.b88345d361482865f1a7af533d41d66e922dcca4c76c2d4b1fcfa65616679471",
		http.MethodPut,
		"/v1/oauth/token",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		404,
		false)
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#Xhca9hTB" // No the ContentType header
	tests++
	_, _, _, err = refreshAccessToken(env,
		"",
		http.MethodPut,
		"/v1/oauth/token",
		"",
		"K7800-H7625-Z5852-N1693-K1972",
		400,
		false)
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#cdWR7DgQ" // No the RequestID Header
	tests++
	_, _, _, err = refreshAccessToken(env,
		"",
		http.MethodPut,
		"/v1/oauth/token",
		"application/json;charset=utf-8",
		"",
		400,
		false)
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#EDOv5oJ6" // Invalid the GET method
	tests++
	_, _, _, err = refreshAccessToken(env,
		"",
		http.MethodGet,
		"/v1/oauth/token",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		405,
		false)
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#Wm1PBdVN" // Invalid the DELETE method
	tests++
	_, _, _, err = refreshAccessToken(env,
		"",
		http.MethodDelete,
		"/v1/oauth/token",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		405,
		false)
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#9CGhxsWz" // OK
	tests++
	_, accessToken, refreshToken, err = refreshAccessToken(env,
		refreshToken,
		http.MethodPut,
		"/v1/oauth/token",
		"application/json;charset=utf-8",
		"K7800-H7625-Z5852-N1693-K1972",
		200,
		true)
	if err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#n4XHQkCf"
	tests++
	if err = getListActiveSessions_200(env, accessToken, publicSessionID); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)
	totalTests = totalTests + tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	tests, err = checklist.revokeRefreshToken(env, accessToken, refreshToken, publicSessionID)
	if err != nil {
		return totalTests, err
	}

	totalTests = totalTests + tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Create a password reset request"
	indentBeforeStatus = "\t\t\t\t\t"
	tests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#fBmC8mR3"
	tests++
	if err = resetUserPassword_400_NoHeaderContentType(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#GQ4XEhtu"
	tests++
	if err = resetUserPassword_400_NoHeaderRequestID(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#M98D7nVP"
	tests++
	if err = resetUserPassword_202(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)
	totalTests = totalTests + tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Confirm password reset (+30 sec of waiting)"
	indentBeforeStatus = "\t\t\t\t"
	tests = 0

	fmt.Print(g.ItemIL)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#V0pUynov"
	tests++
	if linkKey, err = confirmPasswordReset_200(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#k42WCzaL"
	tests++
	if err = confirmPasswordReset_404(env, linkKey); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(tests)), true)
	totalTests = totalTests + tests

	return totalTests, nil
}


