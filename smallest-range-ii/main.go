package main

import (
	"fmt"
	"sort"
)

func main()  {
	a := []int{7, 8, 8}
	k := 5
	fmt.Println(smallestRangeII(a, k))
}

func smallestRangeII(A []int, K int) int {
	if len(A) == 1 {
		return 0
	}
	sort.Ints(A)

	var min = A[len(A) - 1] - A[0]
	e := A[len(A) - 1] - K
	s := A[0] + K
	for ind := 0; ind <= len(A) - 2; ind++ {
		tmp := 0
		if A[ind + 1] - K < s {
			if A[ind] + K > e {
				tmp = A[ind] + K - A[ind + 1] + K
			} else {
				tmp = e - A[ind + 1] + K
			}
		} else {
			if A[ind] + K > e {
				tmp = A[ind] + K - s
			} else {
				tmp = e - s
			}
		}
		if min > tmp {
			min = tmp
		}
	}
	return min
}
