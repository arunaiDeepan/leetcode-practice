func twoSum(nums []int, target int) []int {
    var result []int
	m := make(map[int]int)
	for i, num := range nums {
		dif := target - num
		if anotherIndex, found := m[dif]; found {
			result = append(result, i)
			result = append(result, anotherIndex)
			break
		}
		m[num] = i
	}
	return result
}