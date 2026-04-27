package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.characterReplacement("aabccbb", 2))
	fmt.Println(s.characterReplacement("abbcb", 1))
	fmt.Println(s.characterReplacement("abccde", 1))
}

type Solution struct{}

func (s *Solution) characterReplacement(str string, k int) int {
	maxLength := 0
	windowStart := 0
	maxRepeatLetterCount := 0
	charFrequency := make(map[rune]int)

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := rune(str[windowEnd])
		charFrequency[rightChar]++

		if charFrequency[rightChar] > maxRepeatLetterCount {
			maxRepeatLetterCount = charFrequency[rightChar]
		}

		windowLength := windowEnd - windowStart + 1
		if windowLength-maxRepeatLetterCount > k {
			leftChar := rune(str[windowStart])
			charFrequency[leftChar]--
			windowStart++
		}

		currentLength := windowEnd - windowStart + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
