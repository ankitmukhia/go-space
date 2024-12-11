package refrence

import (
	"strings"
)

func removeProfanity(message *string) {
	myMessage := *message
	myMessage = strings.ReplaceAll(myMessage, "fubb", "****")
	myMessage = strings.ReplaceAll(myMessage, "shiz", "****")
	myMessage = strings.ReplaceAll(myMessage, "witch", "*****")
	*message = myMessage
}

