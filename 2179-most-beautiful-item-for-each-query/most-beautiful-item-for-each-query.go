func maximumBeauty(items [][]int, queries []int) []int {
 	maxInt := math.MaxInt32
	result := [][]int{{0, 0, maxInt}}

	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	for _, item := range items {
		price := item[0]
		beauty := item[1]
		if beauty > result[len(result)-1][1] {
			result[len(result)-1][2] = price
			result = append(result, []int{price, beauty, maxInt})
		}
	}

	var answers []int
	for _, q := range queries {
		for i := len(result) - 1; i >= 0; i-- {
			if result[i][0] <= q {
				answers = append(answers, result[i][1])
				break
			}
		}
	}

	return answers
   
}