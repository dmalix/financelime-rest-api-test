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

	tests, err = checklist.signUp(env)
	if err != nil {
		return totalTests, err
	}

	totalTests += tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	tests, err = checklist.confirmUserEmail(env)
	if err != nil {
		return totalTests, err
	}

	totalTests += tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	newPassword, tests, err = checklist.getPasswordFromEmailMessage(env)
	if err != nil {
		return totalTests, err
	}

	totalTests += tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	accessToken, refreshToken, publicSessionID, tests, err = checklist.requestAccessToken(env, accessToken,
		refreshToken, publicSessionID, newPassword)
	if err != nil {
		return totalTests, err
	}

	totalTests += tests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	tests, err = checklist.getEmailMessageAboutLoginAction(env)
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

	tests, err = checklist.requestUserPasswordReset(env)
	if err != nil {
		return totalTests, err
	}

	totalTests += tests

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
