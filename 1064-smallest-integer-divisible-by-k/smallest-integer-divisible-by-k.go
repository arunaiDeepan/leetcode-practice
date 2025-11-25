func smallestRepunitDivByK(k int) int {
    curr := 1
    for i := 1; i<= k+1; i++ {
        if curr % k == 0 {
            return i
        }
        curr = 10 * (curr % k) + 1
    }
    return -1
}