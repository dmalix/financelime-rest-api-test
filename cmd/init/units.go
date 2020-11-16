package init

import (
	"errors"
	"fmt"
	c "github.com/dmalix/financelime-functional-tests/cmd/config"
	"os/exec"
	"strconv"
	"time"
)

func serviceRestart(env *c.Env) error {

	var (
		cmd      *exec.Cmd
		err      error
		port     = strconv.Itoa(env.Config.SSH.Port)
		instance = fmt.Sprintf("%s@%s", env.Config.SSH.UserName, env.Config.General.Domain)
		unit     = env.Config.Instance.UnitSystemD
	)

	cmd = exec.Command("ssh", "-p", port, instance, "systemctl restart "+unit)
	if err = cmd.Run(); err != nil {
		return errors.New(fmt.Sprintf("Failed to restart service [%s]", err))
	}

	time.Sleep(3 * time.Second)

	cmd = exec.Command("ssh", "-p", port, instance, "systemctl status "+unit)
	if err = cmd.Run(); err != nil {
		return errors.New(fmt.Sprintf("Failed to check of status service [%s]", err))
	}

	return nil
}
