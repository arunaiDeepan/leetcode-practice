func totalMoney(n int) int {
    totalSavings := 0
    weekCount := 0
    day := 0
    for i := 1; i <=n; i++ {
        day = ((i - 1) % 7+ 1)
        totalSavings += weekCount + day

        if i % 7 == 0 {
            weekCount++
        }
    }
    return totalSavings
}