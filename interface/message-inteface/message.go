package message

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
	content := msg.getMessage()
	return content, len(content) * 3
}

type message interface {
	getMessage() string
}

// don't edit below this line
// below 2 getMessage func are 2 different implementation of same interface method signeature
// as we see we dont have to explicitly define any struct to use message interface, Go does it automatically

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}
