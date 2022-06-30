# golang_coin_change

You are given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money.

Return *the fewest number of coins that you need to make up that amount*. If that amount of money cannot be made up by any combination of the coins, return `-1`.

You may assume that you have an infinite number of each kind of coin.

## Examples

**Example 1:**

```
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

```

**Example 2:**

```
Input: coins = [2], amount = 3
Output: -1

```

**Example 3:**

```
Input: coins = [1], amount = 0
Output: 0

```

**Constraints:**

- `1 <= coins.length <= 12`
- `1 <= coins[i] <= 231 - 1`
- `0 <= amount <= 104`

## 解析

題目給一個陣列 coins, 一個整數 amount

coins 代表所有幣值, amount 代表要換的金錢數量

寫一個演算法計算 在可以把 amount 兌換成 coins 所有幣值情下

所需的硬幣最少數量

假設不能兌換則回傳 -1 

思考一個例子：

coins: [1,3,4,5], amount: 7

如果直接用 greedy algorithm 最大幣值優先會算出 用 1 個5元 兩個 1 元來兌換 一共 3 個硬幣

然而實際的最小硬幣數量兌換方式卻是 1 個3元  1個 4 元來兌換 一共 2 個硬幣

透過 決策樹如下

![](https://i.imgur.com/cBn2gPT.png)

會發現當在找尋最佳兌換方式 7 元時

會需要先去解決最佳兌換方式 6 元,最佳兌換方式 4元, 最佳兌換方式 3 元, 最佳兌換方式 2 元 當以各種幣別去兌換時

為了避免重複計算

可以從數量小的量 開始計算最小兌換數量

初始化 兌換 0 元 = 0 因為沒有任何一個幣別可以換

然後逐步往上找 兌換 1 元在所有幣別的最小值

直到找到 兌換 amount 元

對任意兌換 i 元的最小值

初始化 dp[i] = INF

對所有幣別 coin 使得  i - coin ≥ 0 

計算 dp[i] = min(dp[i], 1+dp(1-coin))

算到最後 如果 dp[amount] ≠ INF

代表 dp[amount] 就是所求

因為 對於 每個 amount 每個幣別都要做檢查

所以時間複雜 O(len(coin) * amount)

空間複雜度是 O(amount)

## 程式碼
```go
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

```
## 困難點

1. 不能直接套用 greedy algorithm
2. 需要看出換錢問題的遞迴子關係

## Solve Point

- [x]  建立一個 amount + 1 的陣列 dp 來存放, 每個 amount 的最佳值
- [x]  初始化索引 1 到 amount 為 INF, 還有 dp[0] = 0
- [x]  透過 dp[i] = min(dp[i], 1+ dp[i-coin]), if  i - coin ≥ 0, coin 是 coins 中所有幣值