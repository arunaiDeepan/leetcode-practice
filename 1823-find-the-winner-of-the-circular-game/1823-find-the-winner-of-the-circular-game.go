func findTheWinner(n int, k int) int {
    ans := 0
	for person := 1; person <= n; person++ {
		ans = (ans + k) % person
	}
	return ans + 1
}