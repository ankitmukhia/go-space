package variadicfunc

// check line no: 31 on test file, spread operator is used there.
func sum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum 
}	
