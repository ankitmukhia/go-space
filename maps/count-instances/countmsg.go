package countmsg

func getCounts(messagedUsers []string, validUsers map[string]int) {
	/**
	* I have tought of "collision" here, but it wont happend here. 
	* Firstly map doesn't support duclicate key. and code.
	* secondly code translate, increase count if duplicate find.
	**/
	for _, user := range messagedUsers {
		if _, ok := validUsers[user]; ok {
			validUsers[user]++
		}
	}
}
