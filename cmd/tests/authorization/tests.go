package authorization

import (
	"errors"
	"fmt"
	c "github.com/dmalix/financelime-functional-tests/cmd/config"
	g "github.com/dmalix/financelime-functional-tests/cmd/pgraphics"
	u "github.com/dmalix/financelime-functional-tests/cmd/utils"
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

	accessToken, refreshToken, publicSessionID, tests, err = checklist.requestAccessToken(env, accessToken, refreshToken, publicSessionID, newPassword)
	if err != nil {
		return totalTests, err
	}

	totalTests += tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	tests, err = checklist.getEmailAboutLoginAction(env)
	if err != nil {
		return totalTests, err
	}

	totalTests += tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	tests, err = checklist.getListOfActiveSession(env, accessToken, publicSessionID)
	if err != nil {
		return totalTests, err
	}

	totalTests += tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	accessToken, refreshToken, tests, err = checklist.refreshAccessToken(env, refreshToken, publicSessionID)
	if err != nil {
		return totalTests, err
	}

	totalTests += tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	tests, err = checklist.revokeRefreshToken(env, accessToken, refreshToken, publicSessionID)
	if err != nil {
		return totalTests, err
	}

	totalTests += tests

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
