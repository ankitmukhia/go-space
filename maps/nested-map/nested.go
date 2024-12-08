package nested

func getNameCounts(names []string) map[rune]map[string]int {
	res := make(map[rune]map[string]int)

	for _, name := range names {
		chars := []rune(name) // it will
		for _, chr := range chars{
			/* As for any given outer key(which is here 1st name char) you must check if the inner map exists, and create it if there is not */
			rn, ok := res[chr]
			if !ok {
				rn = make(map[string]int)
				/*
				* '1st chr': map[string]int
				* crete if not exists
				**/
				res[chr] = rn 
			}
			res[chr][name]++
		}
	}

	return res
}
