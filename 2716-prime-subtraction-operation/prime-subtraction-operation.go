func primeSubOperation(nums []int) bool {
	maxElement := getMaxElement(nums)

	sieve := make([]bool, maxElement+1)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[1] = false
	for i := 2; i <= int(math.Sqrt(float64(maxElement+1))); i++ {
		if sieve[i] {
			for j := i * i; j <= maxElement; j += i {
				sieve[j] = false
			}
		}
	}

	currValue := 1
	i := 0
	for i < len(nums) {
		diff := nums[i] - currValue

		if diff < 0 {
			return false
		}

		if sieve[diff] || diff == 0 {
			i++
			currValue++
		} else {
			currValue++
		}
	}
	return true
}

func getMaxElement(nums []int) int {
	max := -1
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}