package best_time_to_buy_and_sell_stock

// by myself
//func maxProfit(prices []int) int {
//	max := 0
//	for i, price := range prices {
//		for _, next := range prices[i+1:] {
//			if next-price > max {
//				max = next - price
//			}
//		}
//	}
//	return max
//}

// 优化 by myself
func maxProfit(prices []int) int {
	i := 0
	j := 1
	max := 0
	for i <= len(prices)-2 && j <= len(prices)-1 {
		if prices[i] > prices[j] {
			i = j
			j = i + 1
		} else {
			if max < prices[j]-prices[i] {
				max = prices[j] - prices[i]
			}
			j++
		}
	}
	return max
}

// high resolution
func maxProfit2(prices []int) int {
	//length := len(prices)
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	min := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > res {
			res = prices[i] - min
		}
	}
	return res
}

// 继续优化
func maxProfit3(prices []int) int {
	//length := len(prices)
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	min := prices[0]
	max := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}
