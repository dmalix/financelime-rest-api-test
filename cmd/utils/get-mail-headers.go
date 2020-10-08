package utils

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"strconv"
)

type IMAP struct {
	Host     string
	Port     int
	Username string
	Password string
}

func (i *IMAP) GetMailHeaders() ([]string, error) {

	var (
		messageIdList []string
		mailboxesList []string
	)

	// Connect to server
	c, err := client.DialTLS(fmt.Sprintf("%s:%s", i.Host, strconv.Itoa(i.Port)), nil)
	if err != nil {
		return messageIdList, err
	}

	// Don't forget to logout
	//noinspection GoUnhandledErrorResult
	defer c.Logout()

	// Login
	if err := c.Login(i.Username, i.Password); err != nil {
		return messageIdList, err
	}

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	for m := range mailboxes {
		mailboxesList = append(mailboxesList, m.Name)
	}

	if err := <-done; err != nil {
		return messageIdList, err
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		return messageIdList, err
	}

	// Get the last 3 messages
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 2 {
		// We're using unsigned integers here, only substract if the result is > 0
		from = mbox.Messages - 2
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	// Get messageId list
	for msg := range messages {
		messageIdList = append(messageIdList, msg.Envelope.MessageId)
	}

	if err := <-done; err != nil {
		return messageIdList, err
	}

	return messageIdList, nil
}
