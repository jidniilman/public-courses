package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.lengthOfLongestSubstringKDistinct("araaci", 2))
	fmt.Println(s.lengthOfLongestSubstringKDistinct("araaci", 1))
	fmt.Println(s.lengthOfLongestSubstringKDistinct("cbbebi", 3))
}

type Solution struct{}

func (s *Solution) lengthOfLongestSubstringKDistinct(str string, k int) int {
	if k == 0 || len(str) == 0 {
		return 0
	}

	maxLength := 0
	windowStart := 0
	charFrequency := make(map[rune]int)

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := rune(str[windowEnd])
		charFrequency[rightChar]++

		for len(charFrequency) > k {
			leftChar := rune(str[windowStart])
			charFrequency[leftChar]--
			if charFrequency[leftChar] == 0 {
				delete(charFrequency, leftChar)
			}
			windowStart++
		}

		currentLength := windowEnd - windowStart + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
