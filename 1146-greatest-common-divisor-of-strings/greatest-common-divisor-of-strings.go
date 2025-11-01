
func gcd(a int, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	lenGcd := gcd(len(str1), len(str2))
	return str1[:lenGcd]    
}