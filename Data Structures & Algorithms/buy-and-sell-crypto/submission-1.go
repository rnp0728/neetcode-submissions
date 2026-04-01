func maxProfit(prices []int) int {
    mp, left := 0, 0

    for right := 1; right < len(prices); right++ {
        profit := prices[right] - prices[left]
        if profit > 0 {
            mp = max(mp, profit)
        } else {
            left = right
        }
    }
    return mp
}
