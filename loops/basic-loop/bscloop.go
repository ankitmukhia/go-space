package bscloop

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		/**
		* 1st: 1 * 0 = 0 // converted to float gives 0.00
		* 2st: 1 * 1 = 1 // converted to float gives 0.01 
		* 3st: 1 * 2 = 2 // converted to float gives 0.02  
		**/
		totalCost += 1.0 + float64(0.01 * float64(i))
	}
	return totalCost 
}
