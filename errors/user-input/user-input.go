package userinput

import "errors"

func validateStatus(status string) error {
	if status == "" {
		err := errors.New("status cannot be empty")	
		return err
	}

	if len(status) > 140 {
		err := errors.New("status exceeds 140 characters")	
		return err
	}

	return nil
}

