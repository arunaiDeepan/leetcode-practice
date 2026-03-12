func repeatedSubstringPattern(s string) bool {
    dup := s+s
    return strings.Contains(dup[1:len(dup)-1], s)   
}