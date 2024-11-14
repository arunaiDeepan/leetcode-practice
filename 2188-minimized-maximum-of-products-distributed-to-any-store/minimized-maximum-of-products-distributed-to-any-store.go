func minimizedMaximum(n int, quantities []int) int {
	left, right := 1, 100000

	for left < right {
		mid := (left + right) / 2
		if f(mid, quantities, n) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func f(x int, quantities []int, n int) bool {
	sum := 0
	for _, a := range quantities {
		sum += (a + x - 1) / x
	}
	return sum > n
}