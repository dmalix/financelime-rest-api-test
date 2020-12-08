package authorization

import (
	"errors"
	"fmt"
	c "github.com/dmalix/financelime-functional-tests/cmd/config"
	g "github.com/dmalix/financelime-functional-tests/cmd/pgraphics"
	u "github.com/dmalix/financelime-functional-tests/cmd/utils"
	"strings"
	"time"
)

func (checkList *checkLists) getPasswordFromEmailMessage(env *c.Env) (string, int, error) {

	const checkListName = "Get a password from the email (+30 sec of waiting)"
	const indentBeforeStatus = "\t\t\t"
	const errorMessage = "Failed to complete the %s test %s [%s]"
	var err error
	var tests = 0
	var testID string
	var newPassword string

	fmt.Print(g.ItemII)
	fmt.Print(checkListName, indentBeforeStatus)

	newPassword, err = test.getPasswordFromEmailMessage(env,
		&u.IMAP{
			Host:     env.Config.IMAP.Host,
			Port:     env.Config.IMAP.Port,
			Username: env.Config.IMAP.Users.Username[1],
			Password: env.Config.IMAP.Users.Password[1],
		})

	testID = "#F2ArtAXB"
	if err != nil {
		return newPassword, tests, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}

	u.Colorize(u.ColorGreen, "Done", true)

	return newPassword, tests, nil
}

func (test *tests) getPasswordFromEmailMessage(env *c.Env, imap *u.IMAP) (string, error) {

	var (
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
