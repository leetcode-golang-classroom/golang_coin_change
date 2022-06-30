package sol

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func coinChange(coins []int, amount int) int {
	minCoins := make([]int, amount+1)
	// minCoins amount -> min coin numbers
	for m := 1; m <= amount; m++ {
		minCoins[m] = math.MaxInt32
		for _, coin := range coins {
			if m-coin >= 0 { // could change
				minCoins[m] = min(minCoins[m], 1+minCoins[m-coin])
			}
		}
	}
	if minCoins[amount] == math.MaxInt32 {
		return -1
	}
	return minCoins[amount]
}
