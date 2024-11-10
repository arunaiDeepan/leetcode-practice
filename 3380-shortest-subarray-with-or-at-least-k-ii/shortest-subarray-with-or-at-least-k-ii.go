func minimumSubarrayLength(nums []int, k int) int {
	count := make([]int, 32)
	start, end, minLength := 0, 0, math.MaxInt32

	for end < len(nums) {
		updateBits(count, nums[end], 1)
		for start <= end && getVal(count) >= k {
			if end-start+1 < minLength {
				minLength = end - start + 1
			}
			updateBits(count, nums[start], -1)
			start++
		}
		end++
	}

	if minLength == math.MaxInt32 {
		return -1
	}
	return minLength
}

func updateBits(count []int, num, val int) {
	for i := 0; i < 32; i++ {
		if (num>>i)&1 == 1 {
			count[i] += val
		}
	}
}

func getVal(count []int) int {
	ans := 0
	for i := 0; i < 32; i++ {
		if count[i] > 0 {
			ans |= (1 << i)
		}
	}
	return ans
}
