package main

import (
	"fmt"
	"strings"
)

func main() {
	s := Solution{}
	fmt.Println(s.largestPalindromic("323211444"))
	fmt.Println(s.largestPalindromic("998877"))
	fmt.Println(s.largestPalindromic("54321"))
}

type Solution struct{}

func (s *Solution) largestPalindromic(num string) string {
	freq := make([]int, 10)
	for _, ch := range num {
		freq[ch-'0']++
	}

	var firstHalf strings.Builder
	middle := ""

	for digit := 9; digit >= 0; digit-- {
		if digit == 0 && firstHalf.Len() == 0 {
			continue
		}

		pairs := freq[digit] / 2
		for i := 0; i < pairs; i++ {
			firstHalf.WriteString(fmt.Sprintf("%d", digit))
		}

		if freq[digit]%2 == 1 && middle == "" {
			middle = fmt.Sprintf("%d", digit)
		}
	}

	if firstHalf.Len() == 0 && middle == "" {
		return "0"
	}

	if firstHalf.Len() == 0 {
		return middle
	}

	firstHalfStr := firstHalf.String()
	secondHalf := s.reverse(firstHalfStr)

	return firstHalfStr + middle + secondHalf
}

func (s *Solution) reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
