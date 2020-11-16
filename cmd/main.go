package main

import (
	"fmt"
	c "github.com/dmalix/financelime-functional-tests/cmd/config"
	i "github.com/dmalix/financelime-functional-tests/cmd/init"
	g "github.com/dmalix/financelime-functional-tests/cmd/pgraphics"
	"github.com/dmalix/financelime-functional-tests/cmd/tests/authorization"
	"github.com/dmalix/financelime-functional-tests/cmd/tests/service"
	u "github.com/dmalix/financelime-functional-tests/cmd/utils"
	"os/exec"
	"strconv"
)

var (
	version   = "unset"
	buildTime = "unset"
	commit    = "unset"
	compiler  = "unset"
)

func main() {

	var (
		err              error
		numberTests      int
		numberTestsTotal int
	)

	env := &c.Env{
		Config: c.Get.Config(),
		Dist:   c.DistData{Version: version, BuildTime: buildTime, Compiler: compiler, Commit: commit},
	}

	if err = confirmRun(env); err != nil {
		return
	}

	u.Colorize(u.ColorBlue, g.Line, true)

	fmt.Print(g.Point0)
	u.Colorize(u.ColorYellow,
		fmt.Sprintf("Functional tests, %s %s [%s, %s]",
			env.Dist.Version, env.Dist.Commit, env.Dist.BuildTime, env.Dist.Compiler), true)

	if env.Config.Instance.InitializationOnStartupEnabled {

		fmt.Println(g.SpaceI)

		if err = i.ServiceInit(env); err != nil {
			u.Colorize(u.ColorRed, "Failed", true)
			u.Colorize(u.ColorRed, g.Line, true)
			u.Colorize(u.ColorRed, err.Error(), true)
			u.Colorize(u.ColorBlue, g.Line, true)
			return
		}
	}

	fmt.Println(g.SpaceI)

	if numberTests, err = service.RunTests(env); err != nil {
		u.Colorize(u.ColorRed, "Failed", true)
		u.Colorize(u.ColorRed, g.Line, true)
		u.Colorize(u.ColorRed, err.Error(), true)
		u.Colorize(u.ColorBlue, g.Line, true)
		return
	}
	numberTestsTotal = numberTestsTotal + numberTests

	fmt.Println(g.SpaceI)

	if numberTests, err = authorization.RunTests(env); err != nil {
		u.Colorize(u.ColorRed, "Failed", true)
		u.Colorize(u.ColorRed, g.Line, true)
		u.Colorize(u.ColorRed, err.Error(), true)
		u.Colorize(u.ColorBlue, g.Line, true)
		return
	}
	numberTestsTotal = numberTestsTotal + numberTests

	fmt.Println(g.SpaceI)
	fmt.Print(g.Point0)

	u.Colorize(u.ColorGreen, fmt.Sprintf("All tests passed successfully(%s)", strconv.Itoa(numberTestsTotal)), true)
	u.Colorize(u.ColorBlue, g.Line, true)
}

func confirmRun(env *c.Env) error {

	var err error
	var out []byte

	out, err = exec.Command("clear").Output()
	if err != nil {
		u.Colorize(u.ColorRed, "Failed", true)
		return err
	}
	u.Colorize(u.ColorYellow, string(out), true)

	out, err = exec.Command("figlet", "Financelime").Output()
	if err != nil {
		u.Colorize(u.ColorRed, "Failed", true)
		return err
	}
	u.Colorize(u.ColorYellow, "\r"+string(out), true)

	fmt.Printf("Are you sure you want to run the REST-API tests for the %s domain?", env.Config.General.Domain)
	if !u.Confirmation() {
		u.Colorize(u.ColorRed, "Start aborted", true)
		return err
	}

	return nil

}
