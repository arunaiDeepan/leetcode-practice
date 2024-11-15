func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)

	left := 0
	for left+1 < n && arr[left] <= arr[left+1] {
		left++
	}

	if left == n-1 {
		return 0
	}

	stack := []int{n - 1}
	for r := n - 2; r > left; r-- {
		if arr[r] <= arr[r+1] {
			stack = append(stack, r)
		} else {
			break
		}
	}

    remove := int(math.Min(float64(n-left-1), float64(stack[len(stack)-1])))

	for i := 0; i <= left; i++ {
		for len(stack) > 0 && arr[i] > arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			remove = int(math.Min(float64(remove), float64(stack[len(stack)-1]-i-1)))
		}
	}

	return remove 
}