package main

func predictTheWinner(nums []int) bool {
	N := len(nums)

	hasCache := make([][]bool, N+1)
	cache := make([][]int, N+1)
	for i := range hasCache {
		hasCache[i] = make([]bool, N+1)
		cache[i] = make([]int, N+1)
	}

	var score func(int, int) int
	score = func(left, right int) int {
		if left > right {
			return 0
		}

		if hasCache[left][right] {
			return cache[left][right]
		}

		scoreTakingLeft := nums[left] - score(left+1, right)
		scoreTakingRight := nums[right] - score(left, right-1)

		hasCache[left][right] = true
		cache[left][right] = max(scoreTakingLeft, scoreTakingRight)
		return cache[left][right]
	}

	return score(0, N-1) >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
