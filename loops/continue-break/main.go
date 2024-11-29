package main

import "fmt"

func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}

		if n % 2 == 0 {
			continue
		}
		// squaring number means multiplying the number be itself. And result is called square.
		isPrime := true
		for i := 3; i * i <= n; i+=2 {
			if n % i == 0 {
				isPrime = false
				break	
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}

