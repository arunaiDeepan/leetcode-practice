func threeConsecutiveOdds(arr []int) bool {
    c := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			c = 0
		}
		if arr[i]%2 != 0 {
			c += 1
			if c < 3 {
				continue
			}
		}
		if c == 3 {
			return true
		}

	}
	return false
}