package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {

	const configFileName string = "../config/config.json"

	const ( // Confidential Settings
		envInviteCode           = "FL_TEST_INVITE_CODE"
		envSSHPort              = "FL_TEST_SSH_PORT"
		envSSHUserName          = "FL_TEST_SSH_USERNAME"
		envImapUsername         = "FL_TEST_IMAP_USERNAME"
		envImapPassword         = "FL_TEST_IMAP_PASSWORD"
		envImapHost             = "FL_TEST_IMAP_HOST"
		envCurrentRemoteAddress = "FL_TEST_CURRENT_REMOTE_ADDRESS"
	)

	var (
		err  error
		body []byte
	)

	body, err = ioutil.ReadFile(configFileName)
	if err != nil {
		log.Fatalf("Failed to read %s file [%s]", configFileName, err.Error())
		return
	}

	err = json.Unmarshal(body, &config)
	if err != nil {
		log.Fatalf("Failed to convert %s file to struct [%s]", configFileName, err.Error())
		return
	}

	config.InviteCode = os.Getenv(envInviteCode)
	config.SSH.Port, err = strconv.Atoi(os.Getenv(envSSHPort))
	if err != nil {
		log.Fatalf("Failed to convert envSSHPort value to int format [%s]", err.Error())
		return
	}
	config.SSH.UserName = os.Getenv(envSSHUserName)
	config.IMAP.Users.Username = strings.Split(os.Getenv(envImapUsername), " ")
	config.IMAP.Users.Password = strings.Split(os.Getenv(envImapPassword), " ")
	config.IMAP.Host = os.Getenv(envImapHost)
	config.CurrentRemoteAddress = os.Getenv(envCurrentRemoteAddress)
}
