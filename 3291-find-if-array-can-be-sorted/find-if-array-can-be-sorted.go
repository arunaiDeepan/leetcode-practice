func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canSortArray(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}

	pmax := 0
	cmin, cmax := 0, 0
	pcnt := 0

	for _, v := range nums {
		ccnt := bits.OnesCount(uint(v))
		if pcnt == ccnt {
			cmin = min(cmin, v)
			cmax = max(cmax, v)
		} else if cmin < pmax {
			return false
		} else {
			pmax = cmax
			cmin, cmax = v, v
			pcnt = ccnt
		}
	}

	return cmin >= pmax
}