package main

//  Two solutions

/* func getLast[T any](s []T) T {
	last := len(s) - 1
	var myZero T

	for i := 0; i < len(s); i++ {
		myZero = s[last]
	}
		
	return myZero 
} */

func getLast[T any](s []T) T {
	var myZero T
	if len(s) == 0 {
		return myZero
	}

	return s[len(s)-1]
}

