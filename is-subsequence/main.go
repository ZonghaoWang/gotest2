package main

import "fmt"

func isSubsequence(s string, t string) bool {
	sl := len(s)
	tl := len(t)
	lastExists := make([]int, sl)
	currentExists := make([]int, sl)
	if tl < sl {
		return  false
	}
	if sl == 0 {
		return true
	}
	if s[0] == t[0] {
		lastExists[0] = 1
	}
	for index := 1; index < tl; index ++ {
		for ind := 0; ind < sl; ind ++ {
			sc := s[ind]
			tc := t[index]
			if sc == tc {
				if ind > 0 {
					currentExists[ind] = lastExists[ind - 1] + lastExists[ind]
				} else {
					currentExists[ind] = lastExists[ind] + 1
				}
			} else {
				currentExists[ind] = lastExists[ind]
			}
		}
		if currentExists[sl - 1] > 0 {
			return true
		}
		copy(lastExists, currentExists)
	}
	return false
}
func main()  {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
