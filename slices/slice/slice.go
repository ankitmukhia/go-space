package slice

import "errors"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planPro {
		return messages[:], nil
	}

	if plan == planFree {
		return messages[0:2], nil
	}
	// nil fulfilles []string. It represents "no value || not initialized", for 
	// pointers, slices, maps,  channels, and function. so in our case it is slices
	return nil, errors.New("unsupported plan") 
}

