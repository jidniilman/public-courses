package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.minWindow("aabdec", "abc"))
	fmt.Println(s.minWindow("aabdec", "abac"))
	fmt.Println(s.minWindow("abdbca", "abc"))
	fmt.Println(s.minWindow("adcad", "abc"))
}

type Solution struct{}

func (s *Solution) minWindow(str string, pattern string) string {
	patternFreq := make(map[rune]int)
	for _, ch := range pattern {
		patternFreq[ch]++
	}

	windowStart := 0
	matched := 0
	minLength := len(str) + 1
	substrStart := 0

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := rune(str[windowEnd])

		if _, exists := patternFreq[rightChar]; exists {
			patternFreq[rightChar]--
			if patternFreq[rightChar] >= 0 {
				matched++
			}
		}

		for matched == len(pattern) {
			currentLength := windowEnd - windowStart + 1
			if currentLength < minLength {
				minLength = currentLength
				substrStart = windowStart
			}

			leftChar := rune(str[windowStart])
			windowStart++

			if _, exists := patternFreq[leftChar]; exists {
				if patternFreq[leftChar] == 0 {
					matched--
				}
				patternFreq[leftChar]++
			}
		}
	}

	if minLength > len(str) {
		return ""
	}
	return str[substrStart : substrStart+minLength]
}
