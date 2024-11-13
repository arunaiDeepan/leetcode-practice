
func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var result int64

	for i := 0; i < len(nums)-1; i++ {
		low := lowerBound(nums, i+1, len(nums), lower-nums[i])
		up := upperBound(nums, i+1, len(nums), upper-nums[i])
		result += int64(up - low)
	}

	return result
}

func lowerBound(nums []int, start, end, target int) int {
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] >= target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

func upperBound(nums []int, start, end, target int) int {
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}