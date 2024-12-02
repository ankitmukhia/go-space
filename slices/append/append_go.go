package append

type cost struct {
	day int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	daycost := []float64{}		
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(daycost) {
			daycost = append(daycost, 0.0)	
		}

		daycost[cost.day] += cost.value
	}

	return daycost
}
