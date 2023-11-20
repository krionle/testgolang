package main

import "fmt"

func main() {
	s := []int{1, 123, 2, 312, 31, 31312}

	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				s1 := s[j]
				s2 := s[j+1]
				s[j] = s2
				s[j+1] = s1
			}
		}
	}
	for i, v := range s {
		fmt.Printf("%d:%d\n", i, v)
	}
}
