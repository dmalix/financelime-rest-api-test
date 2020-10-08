package init

import (
	"errors"
	"fmt"
	c "github.com/dmalix/financelime-rest-api-test/cmd/config"
	g "github.com/dmalix/financelime-rest-api-test/cmd/pgraphics"
	u "github.com/dmalix/financelime-rest-api-test/cmd/utils"
)

func ServiceInit(env *c.Env) error {

	var (
		err                error
		taskName           string
		indentBeforeStatus string
		taskID             string
	)

	fmt.Print(g.PointI0)
	u.Colorize(u.ColorYellow, "Initialization", true)
	fmt.Println(g.SpaceII)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	taskName = "The REST-API service restart"
	indentBeforeStatus = "\t\t\t\t\t\t"

	fmt.Print(g.ItemIL)
	fmt.Print(taskName, indentBeforeStatus)

	taskID = "#UE17kj9c"
	if err = serviceRestart(env); err != nil {
		return errors.New(fmt.Sprintf("Failed to restart service %s [%s]", taskID, err))
	}

	u.Colorize(u.ColorGreen, "Done", true)

	return nil
}
