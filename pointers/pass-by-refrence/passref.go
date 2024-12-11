package passref


type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}


func getMessageText(analytics *Analytics, message Message) {
	recipient := message.Recipient
	success := message.Success

	if success {
		analytics.MessagesTotal++
		analytics.MessagesSucceeded++
	} else {
		if recipient != "" {
			analytics.MessagesTotal++
		}
		analytics.MessagesFailed++
	}
}
