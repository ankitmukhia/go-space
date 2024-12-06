package tagger

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i, msg := range messages {
		// returns : ['Urgent', 'Promo']
		result := tagger(msg)

		messages[i].tags = result
	}
	return messages 
}

// expected: [][]string{{"Urgent"}, {"Promo"}},

func tagger(msg sms) []string {
	tags := []string{}
	/*
	* msg contain : id: '', content: ''
	*/

	firstTag := strings.Contains(strings.ToLower(msg.content), "urgent")	
	secondTag := strings.Contains(strings.ToLower(msg.content), "sale")

	if firstTag {
		tags = append(tags, "Urgent")
	} 

	if secondTag {
		tags = append(tags, "Promo")
	}

	return tags
}
