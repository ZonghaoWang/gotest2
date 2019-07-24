package main

import "fmt"

func numDistinct(s string, t string) int {
	sl := len(s)
	tl := len(t)
	rst := make([][]int, tl)
	for i := 0; i < tl; i++ {
		rst[i] = make([]int, sl)
	}
	if s[0] == t[0] {
		rst[0][0] = 1
	}
	for si := 1; si < sl; si++ {
		for ti := 0; ti < tl; ti++ {
			tc := t[ti]
			sc := s[si]
			if sc == tc {
				if ti > 0 {
					rst[ti][si] = rst[ti - 1][si - 1] + rst[ti][si - 1]
				} else {
					rst[ti][si] = rst[ti][si - 1] + 1
				}
			} else {
					rst[ti][si] = rst[ti][si - 1]
			}
		}
	}
	return rst[tl - 1][sl - 1]
}

func main() {
	s := "rabbbit"
	t := "rabit"
	fmt.Println(numDistinct(s, t))
}