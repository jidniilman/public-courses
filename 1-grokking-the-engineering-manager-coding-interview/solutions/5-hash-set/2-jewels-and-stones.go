package main

import "fmt"

func main() {
	s := Solution{}
	fmt.Println(s.numJewelsInStones("abc", "aabbcc"))
	fmt.Println(s.numJewelsInStones("aA", "aAaZz"))
	fmt.Println(s.numJewelsInStones("zZ", "zZzZzZ"))
}

type Solution struct{}

func (s *Solution) numJewelsInStones(jewels string, stones string) int {
	jewelSet := make(map[rune]bool)
	for _, jewel := range jewels {
		jewelSet[jewel] = true
	}

	count := 0
	for _, stone := range stones {
		if jewelSet[stone] {
			count++
		}
	}

	return count
}
