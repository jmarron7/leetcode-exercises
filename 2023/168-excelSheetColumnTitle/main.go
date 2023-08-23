package main

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) <= 2 {
		return len(s)
	}
	left, right := 0, 0
	charMap := make(map[byte]int)
	maxLength := 2

	for right < len(s) {
		if len(charMap) <= 2 {
			charMap[s[right]] = right
			right++
		}

		if len(charMap) > 2 {
			leftIndex := len(s)
			for _, index := range charMap {
				if index < leftIndex {
					leftIndex = index
				}
			}

			delete(charMap, s[leftIndex])
			left = leftIndex + 1
		}

		maxLength = max(maxLength, right-left)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))   // Output: 3
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb")) // Output: 5
}
