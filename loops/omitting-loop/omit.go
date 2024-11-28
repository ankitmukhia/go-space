package omit

func maxMessages(thresh int) int {
	// loops forever until it exede our thrash(max number), until then
	// it bruteforces.
	max := 0
	for i := 0; ; i++ {
		/**
		* max + 100 + loop index
		**/
		max += 100 + (1 * i)
		if max > thresh {
			return i
		}
	}

	return max
}


