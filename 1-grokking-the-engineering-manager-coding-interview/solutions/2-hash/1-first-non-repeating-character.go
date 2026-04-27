package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.firstUniqChar("abab"))
}

type Solution struct{}

func (s *Solution) firstUniqChar(sStr string) int {
	frequencyMap := map[string]int{}
	// Trace and count the string characters
	for _, r := range sStr {
		frequencyMap[string(r)]++
	}

	// Find the first non-repeating character index or return -1
	for i, r := range sStr {
		if frequencyMap[string(r)] == 1 {
			return i
		}
	}

	return -1
}
