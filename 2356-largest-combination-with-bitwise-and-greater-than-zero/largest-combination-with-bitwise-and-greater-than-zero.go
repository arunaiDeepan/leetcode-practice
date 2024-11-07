func largestCombination(candidates []int) int {
	result := 0
	for i := 0; i < 32; i++ {
		count := 0
		for _, candidate := range candidates {
			if candidate&(1<<i) != 0 {
				count++
			}
		}
		if count > result {
			result = count
		}
	}
	return result  
}