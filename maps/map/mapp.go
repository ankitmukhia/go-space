package mapp

import (
	"errors"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	result := make(map[string]user)

	for i, name := range names {
		result[name] = user{
			name: name,
			phoneNumber: phoneNumbers[i],
		}
	}

	return result, nil
}

type user struct {
	name        string
	phoneNumber int
}

