package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.minSteps("designgurus", "garumdespgn"))
	fmt.Println(s.minSteps("abc", "def"))
	fmt.Println(s.minSteps("listen", "silent"))
}

type Solution struct{}

func (s *Solution) minSteps(sStr string, tStr string) int {
	charFreq := map[rune]int{}

	for _, char := range sStr {
		charFreq[char]++
	}

	for _, char := range tStr {
		charFreq[char]--
	}

	steps := 0
	for _, count := range charFreq {
		if count > 0 {
			steps += count
		}
	}

	return steps
}
