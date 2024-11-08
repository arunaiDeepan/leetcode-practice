func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	xorr := nums[0]
	maxXor := (1 << maximumBit) - 1

	for i := 1; i < n; i++ {
		xorr ^= nums[i]
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = xorr ^ maxXor
		xorr ^= nums[n-1-i]
	}

	return ans
}