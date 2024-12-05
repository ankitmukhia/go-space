package rangee

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, message := range msg {
		for _, badWord := range badWords {
			if message == badWord {
			  return i
			}
		}
	}

	return -1
}
