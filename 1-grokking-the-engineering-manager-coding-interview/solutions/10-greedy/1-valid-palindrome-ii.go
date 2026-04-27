package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.validPalindrome("racecar"))
	fmt.Println(s.validPalindrome("abccdba"))
	fmt.Println(s.validPalindrome("abcdef"))
}

type Solution struct{}

func (s *Solution) validPalindrome(str string) bool {
	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return s.isPalindrome(str, left+1, right) || s.isPalindrome(str, left, right-1)
		}
		left++
		right--
	}

	return true
}

func (s *Solution) isPalindrome(str string, left, right int) bool {
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
