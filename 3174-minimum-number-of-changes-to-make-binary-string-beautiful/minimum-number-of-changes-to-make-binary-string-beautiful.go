func minChanges(s string) int {
 	n := len(s)
	count := 0
	for i := 0; i < n-1; i++ {
		if s[i] != s[i+1] {
			count++
		}
        
		i++
	}
	return count   
}