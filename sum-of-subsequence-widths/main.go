package main

import (
	"fmt"
)

func main()  {
	fmt.Println(sumSubseqWidths([]int{3,7,2,3}))
}

func intMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func intMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func sumSubseqWidths(A []int) int {
	mp := make(map[int]int, len(A))
	for _, i := range A {
		mp[i]++
	}
	var rst int
	var baseMod int = 1000000009
	for m := range A {
		mm := A[m]
		mp[mm]++
		max, min := m, m
		tmp := 0
		ori := 0
		for n := m + 1; n <  len(A); n++ {
			nm := A[n]
			if len(mp) == 1{
				if nm == mm {
					mp[nm]++
				} else {
					max = intMax(max, nm)
					min = intMin(min, nm)
					ori = max - min
					tmp += ori
				}
			} else {
				if nm == max || nm == min {
					ori = ori / mp[nm] * (mp[nm] + 1)
				} else {

				}
			}

		}
	}
	return int(rst % 1000000009)
}