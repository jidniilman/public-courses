package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.closeStrings("aacbbc", "bbcaca"))
	fmt.Println(s.closeStrings("xxyyzz", "zzxxyy"))
	fmt.Println(s.closeStrings("aabbcc", "aabbc"))
}

type Solution struct{}

func (s *Solution) sortInts(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func (s *Solution) closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	freq1 := make(map[rune]int)
	freq2 := make(map[rune]int)

	for _, char := range word1 {
		freq1[char]++
	}

	for _, char := range word2 {
		freq2[char]++
	}

	if len(freq1) != len(freq2) {
		return false
	}

	for char := range freq1 {
		if _, exists := freq2[char]; !exists {
			return false
		}
	}

	counts1 := []int{}
	counts2 := []int{}

	for _, count := range freq1 {
		counts1 = append(counts1, count)
	}

	for _, count := range freq2 {
		counts2 = append(counts2, count)
	}

	s.sortInts(counts1)
	s.sortInts(counts2)

	for i := 0; i < len(counts1); i++ {
		if counts1[i] != counts2[i] {
			return false
		}
	}

	return true
}
