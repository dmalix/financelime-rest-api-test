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

func (checkList *checkLists) getEmailMessageAboutLoginAction(env *c.Env) (int, error) {

	const checkListName = "Get an email about a login action (+30 sec of waiting)"
	const indentBeforeStatus = "\t\t"
	const errorMessage = "Failed to complete the %s test %s [%s]"
	var err error
	var tests = 0
	var testID string

	fmt.Print(g.ItemII)
	fmt.Print(checkListName, indentBeforeStatus)

	testID = "#piQaaBr0" // OK
	if err = test.getEmailMessageAboutLoginAction(env); err != nil {
		return 0, errors.New(fmt.Sprintf(errorMessage, checkListName, testID, err))
	}

	u.Colorize(u.ColorGreen, "Done", true)

	return tests, nil
}

func (test *tests) getEmailMessageAboutLoginAction(env *c.Env) error {

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
