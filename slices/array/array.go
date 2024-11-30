package array

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	arrstr := [3]string{primary, secondary, tertiary}
	arrint := [3]int{len(primary), len(secondary) + len(primary), len(tertiary) + len(primary) + len(secondary)}

	return arrstr, arrint
}	
