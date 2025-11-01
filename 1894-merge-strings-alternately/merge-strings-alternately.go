func mergeAlternately(word1 string, word2 string) string {
	lenWord1 := len(word1)
	lenWord2 := len(word2)
	maxLenght := max(lenWord1, lenWord2)

	result := make([]byte, 0, lenWord1+lenWord2)

	for i := 0; i < maxLenght; i++ {
		if i < lenWord1 {
			result = append(result, word1[i])
		}
		if i < lenWord2 {
			result = append(result, word2[i])
		}
	}

	return string(result)  
}