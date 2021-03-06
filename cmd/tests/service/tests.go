package service

import (
	"errors"
	"fmt"
	c "github.com/dmalix/financelime-functional-tests/cmd/config"
	g "github.com/dmalix/financelime-functional-tests/cmd/pgraphics"
	u "github.com/dmalix/financelime-functional-tests/cmd/utils"
	"strconv"
)

func RunTests(env *c.Env) (int, error) {

	var (
		err                error
		testName           string
		testID             string
		indentBeforeStatus string
		numberTests        int
		numberTestsTotal   int
	)
	const errorMessage = "Failed to complete the test %s [%s]"

	fmt.Print(g.PointI0)
	u.Colorize(u.ColorYellow, "Service tests group", true)
	fmt.Println(g.SpaceII)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	testName = "Get the current version of the REST-API service"
	indentBeforeStatus = "\t\t\t"
	numberTests = 0

	fmt.Print(g.ItemIL)
	fmt.Print(testName, indentBeforeStatus)

	testID = "#PV6ClE5y"
	numberTests++
	if err = getCurrentVersion_400_NoHeaderRequestID(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#fnK0j9K2"
	numberTests++
	if err = getCurrentVersion_400_InvalidRequestID(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	testID = "#EPV6Cl5y"
	numberTests++
	if err = getCurrentVersion_200(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, testID, err))
	}

	u.Colorize(u.ColorGreen, fmt.Sprintf("Pass(%s)", strconv.Itoa(numberTests)), true)
	numberTestsTotal = numberTestsTotal + numberTests

	return numberTestsTotal, nil
}
