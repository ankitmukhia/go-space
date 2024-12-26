package main

func concurrentFib(n int) []int {
	val := []int{}		
	s := make(chan int)

	go fibonacci(n, s)
	for item := range s {
		val = append(val, item)
	}
	return val
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
