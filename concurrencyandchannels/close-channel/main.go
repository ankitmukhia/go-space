package main

/* https://www.boot.dev/lessons/bb1f32af-1ca2-4d6e-9bb1-226d9bb1232b */

func countReports(numSentCh chan int) int {
	// my initial solu: val, ok := <-numSentCh
	// which is wrong coz, it will only run once when func calls, so my ok will remain true
	num := 0

	for {
		val, ok := <-numSentCh
		if !ok {
			break
		}
		num += val		
	}
	return num 
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

