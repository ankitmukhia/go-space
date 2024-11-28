package nowhileloop

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	// initial const of the sending one message in pennies
	actualCostInPennies := 1.0
	// initilize the num of message that can be sent
	maxMessageSended := 1

	// calculate the initial balance after the first message
	balance := float64(maxCostInPennies) - actualCostInPennies // balance = 5 - 1 = 4

	// after fourth iteration our balance 

	// you’re mentally thinking about when you can’t send another message anymore:
	// But in programming, the loop condition specifies when it should keep running, not when it should stop. 
	// use !(balance < actualCostInpennies) if it feels confuinsg

	for balance >= actualCostInPennies { 		
		// First Iterations: 
		actualCostInPennies *= costMultiplier // const of the next message = (1.0 * 1.1  = 1.1)

		balance -= actualCostInPennies // Remaining balance (4 - 1.1 = 2.9)

		maxMessageSended++ // increment the count of messages = 2

		// Second Iterations:	
		// actualCostInPennies (1.1 * 1.1 = 1.21) 
		// balance (2.9 - 1.21 = 1.69)
		// maxMessageSended = 3


		// Third Iterations:	
		// actualCostInPennies (1.21 * 1.1 = 1.331) 
		// balance (1.69 - 1.331 = 0.359)
		// maxMessageSended = 4 

		// fourth Iterations: STOP:: STOP:: STOP::	
		// actualCostInPennies (1.331 * 1.1 = 1.4641)  
		// balance (0.359 - 1.4641 = -1.1051) // here balance is 0.359 and my message consts 1.4641 :: so it will stop because my balance is 
		// maxMessageSended = 5 
	}

	if balance < 0 {
		maxMessageSended--
	}

	return maxMessageSended
}
