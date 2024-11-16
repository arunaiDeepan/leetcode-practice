func resultsArray(nums []int, k int) []int {
    n := len(nums)
	if n == 1 || k == 1 {
		return nums
	}

	result := make([]int, n-k+1)
	for i := range result {
		result[i] = -1
	}

	length := 1

	for r := 1; r < n; r++ {
		if nums[r] == nums[r-1]+1 {
			length++
		} else {
			length = 1
		}

		if length >= k {
			result[r-k+1] = nums[r]
		}
	}

	return result
}