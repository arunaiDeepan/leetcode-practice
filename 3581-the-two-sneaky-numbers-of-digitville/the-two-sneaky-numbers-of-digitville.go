func getSneakyNumbers(nums []int) []int {
	result := []int{}
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
		if freq[num] == 2 {
			result = append(result, num)
		}
	}
	return result    
}