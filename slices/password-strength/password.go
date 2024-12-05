package password 

import "unicode"

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}

	isUpper := false
	isDigit := false

	for _, pss := range password {
		/** 
		* Wrong :: coz this check, whether isUpper & isDigit are true in the same character
		* if unicode.IsUpper(pss) && unicode.IsDigit(pss)
		**/

		if unicode.IsUpper(pss) {
			isUpper = true
		}
		if unicode.IsDigit(pss) {
			isDigit = true
		}

		// if true then return dont check further
		if isUpper && isDigit {
			return true
		}
	}

	return isUpper && isDigit
}
