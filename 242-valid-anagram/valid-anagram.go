func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }

    freq := make(map[rune]int)

    for _ , c := range s {
        freq[c]++
    }
    
    for _ , c := range t {
        if _ , exists := freq[c]; !exists {
            return false
        }

        freq[c] -= 1

        if freq[c] < 0 {
            return false
        }
    }

    return true
}
