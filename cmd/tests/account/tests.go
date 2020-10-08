package account

import (
	"errors"
	"fmt"
	c "github.com/dmalix/financelime-rest-api-test/cmd/config"
	g "github.com/dmalix/financelime-rest-api-test/cmd/pgraphics"
	u "github.com/dmalix/financelime-rest-api-test/cmd/utils"
	"strconv"
)

func RunTests(env *c.Env) (int, error) {

	var (
		err                error
		testName           string
		numberTests        int
		numberTestsTotal   int
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
	u.Colorize(u.ColorYellow, "Account tests group", true)
	fmt.Println(g.SpaceII)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Create a new account"
	indentBeforeStatus = "\t\t\t\t\t\t\t"
	numberTests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#hdpFBOG7"
	numberTests++
	if err = createNewAccount_400_NoHeaderRequestID(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#VuLq72Fi"
	numberTests++
	if err = createNewAccount_400_NoHeaderContentType(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#T6PwbuNa"
	numberTests++
	if err = createNewAccount_400_InvalidLanguageDe(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#02GXjc2C"
	numberTests++
	if err = createNewAccount_400_InvalidLanguage1234(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#yBo3E2X4"
	numberTests++
	if err = createNewAccount_400_InvalidLanguage(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#9C73chj1"
	numberTests++
	if err = createNewAccount_400_InvalidEmail(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#tc9Id2BU"
	numberTests++
	if err = createNewAccount_409_InvalidInviteCode(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#VDR1Hmzv"
	numberTests++
	if err = createNewAccount_409_AccountAlreadyExist0(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#nRIxV62A"
	numberTests++
	if err = createNewAccount_202_Email1(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#dREV5jQx"
	numberTests++
	if err = createNewAccount_409_AccountAlreadyExist1(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#dREV5jQx"
	numberTests++
	if err = createNewAccount_409_AccountAlreadyExist1(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#Dyr5SSNC"
	numberTests++
	if err = createNewAccount_202_Email2(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#DwStF6M0"
	numberTests++
	if err = createNewAccount_409_AccountAlreadyExist2(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(numberTests)), true)
	numberTestsTotal = numberTestsTotal + numberTests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Confirm email (+30 sec of waiting)"
	indentBeforeStatus = "\t\t\t\t\t"
	numberTests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#G7zv5jx4"
	numberTests++
	if err = confirmEmail_404_InvalidEndPoint123456789(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#1iSKofF8"
	numberTests++
	if linkKey, err = confirmEmail_200(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#eJ7RFkI5"
	numberTests++
	if err = confirmEmail_404(env, linkKey); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#DwStF6M0"
	numberTests++
	if err = createNewAccount_409_AccountAlreadyExist2(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(numberTests)), true)
	numberTestsTotal = numberTestsTotal + numberTests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Get a password from the email (+30 sec of waiting)"
	indentBeforeStatus = "\t\t\t"
	numberTests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#F2ArtAXB"
	if newPassword, err = getPasswordFromEmail(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, "Done", true)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Get Access Token"
	indentBeforeStatus = "\t\t\t\t\t\t\t"
	numberTests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#Kp0gJ1mm"
	numberTests++
	if accessToken, refreshToken, err = getAccessToken_200(env, newPassword); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(numberTests)), true)
	numberTestsTotal = numberTestsTotal + numberTests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Get an email about a login action (+30 sec of waiting)"
	indentBeforeStatus = "\t\t"
	numberTests = 0

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
	numberTests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#MOi0O3GS"
	numberTests++
	if err = getListActiveSessions_403_NoHeaderAuthorization(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#jg7U4TDR"
	numberTests++
	if err = getListActiveSessions_403_InvalidAccessToken(env, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJGaW5hbmNlbGltZS5jb20iLCJzdWIiOiJBdXRob3JpemF0aW9uIiwicHVycG9zZSI6ImFjY2VzcyIsImlkIjoiMTZkN2RiNTM3MjQ3ZWFmMTEzZjRjOGFkNTllOWEyYTU4OWNlN2NhZjYxMzViY2Q3YmZlYzBiNTI1YWY0OGEyYSIsImlhdCI6MTU5NjgzNTM5OX0.d68bea3232f10c60483f838fff8d8c66661cb497b141213c9a006be2e7c9d723"); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#C9a10Hr1"
	numberTests++
	if err = getListActiveSessions_403_NoBearerValueIntoHeaderAuthorization(env, accessToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#VaK6YsxC"
	numberTests++
	if err = getListActiveSessions_400_NoHeaderRequestID(env, accessToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#6R5YXIzJ"
	numberTests++
	if err = getListActiveSessions_405_InvalidMethodPost(env, accessToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#KnGUxH36"
	numberTests++
	if err = getListActiveSessions_200(env, accessToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(numberTests)), true)
	numberTestsTotal = numberTestsTotal + numberTests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Refresh Access Token"
	indentBeforeStatus = "\t\t\t\t\t\t\t"
	numberTests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#pWs8HJ4F"
	numberTests++
	if err = refreshAccessToken_403_InvalidRefreshToken(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#Xhca9hTB"
	numberTests++
	if err = refreshAccessToken_400_NoHeaderContentType(env, refreshToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#cdWR7DgQ"
	numberTests++
	if err = refreshAccessToken_400_NoHeaderRequestID(env, refreshToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#EDOv5oJ6"
	numberTests++
	if err = refreshAccessToken_400_InvalidMethodPost(env, refreshToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#Wm1PBdVN"
	numberTests++
	if err = refreshAccessToken_405_InvalidMethodGet(env, refreshToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#9CGhxsWz"
	numberTests++
	if publicSessionID, accessToken, refreshToken, err = refreshAccessToken_200(env, refreshToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#n4XHQkCf"
	numberTests++
	if err = getListActiveSessions_200(env, accessToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(numberTests)), true)
	numberTestsTotal = numberTestsTotal + numberTests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Revoke Access Token"
	indentBeforeStatus = "\t\t\t\t\t\t\t"
	numberTests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#M678F5tr"
	numberTests++
	if err = revokeAccessToken_400_NoHeaderRequestID(env, accessToken, publicSessionID); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#2NKyviy2"
	numberTests++
	if err = revokeAccessToken_400_NoHeaderContentType(env, accessToken, publicSessionID); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#Wp4oRz1n"
	numberTests++
	if err = revokeAccessToken_405_InvalidMethodPut(env, accessToken, publicSessionID); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#wz0FKQHO"
	numberTests++
	if err = revokeAccessToken_405_InvalidMethodGet(env, accessToken, publicSessionID); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#Gq7dHBlM"
	numberTests++
	if err = revokeAccessToken_200(env, accessToken, publicSessionID); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#7JGP0R3o"
	numberTests++
	if err = getListActiveSessions_403_InvalidAccessToken(env, accessToken); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(numberTests)), true)
	numberTestsTotal = numberTestsTotal + numberTests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Create a password reset request"
	indentBeforeStatus = "\t\t\t\t\t"
	numberTests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#fBmC8mR3"
	numberTests++
	if err = resetAccountPassword_400_NoHeaderContentType(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#GQ4XEhtu"
	numberTests++
	if err = resetAccountPassword_400_NoHeaderRequestID(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#M98D7nVP"
	numberTests++
	if err = resetAccountPassword_202(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(numberTests)), true)
	numberTestsTotal = numberTestsTotal + numberTests

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Confirm password reset (+30 sec of waiting)"
	indentBeforeStatus = "\t\t\t\t"
	numberTests = 0

	fmt.Print(g.ItemII)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#V0pUynov"
	numberTests++
	if linkKey, err = confirmPasswordReset_200(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#k42WCzaL"
	numberTests++
	if err = confirmPasswordReset_404(env, linkKey); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(numberTests)), true)
	numberTestsTotal = numberTestsTotal + numberTests

	return numberTestsTotal, nil
}
