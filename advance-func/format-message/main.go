package main

import "fmt"

func add(x, y int) int {
	return x + y 
}

func mul(x, y int) int {
	return x * y 
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	/**
	* outer arithmetic func, also takes two argunment, (return value of innter arithmetic, c) 
	* so inner arithmetic will return 2 + 3 = 5 & add the with c(4) and bcom 9. Same mult.
	**/
	return arithmetic(arithmetic(a, b), c)
}

func main() {
	fmt.Println(aggregate(2,3,4, add))
	fmt.Println(aggregate(2,3,4, mul))
}
