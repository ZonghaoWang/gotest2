package main

import (
	"fmt"
	"sort"
)

func main()  {
	a := [][]int{[]int{1,3}, []int{6, 9}}
	b := []int{2, 5}
	fmt.Println(insert(a, b))
}

func merge(a, b []int) []int {
	var l, r int
	if b[0] >= a[0] && b[0] <= a[1] {
		l = a[0]
		if a[1] > b[1] {
			r = a[1]
		} else {
			r = b[1]
		}
	} else {
		if b[0] > a[1] {
			return nil
		} else {
			l = b[0]
			if b[1] < a[0] {
				return nil
			} else {
				if b[1] > a[1] {
					r = b[1]
				} else {
					r = a[1]
				}
			}
		}
	}
	return []int{l, r}
}

func insert(intervals [][]int, newInterval []int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	hit := 0
	rst := make([][]int, 0, len(intervals) + 1)
	rst = append(rst, newInterval)
	for _, l := range intervals {
		tmp := rst[len(rst) - 1]
		rst = rst[0: len(rst) - 1]
		if s := merge(l, tmp); s != nil {
			rst = append(rst, s)
			hit++
		} else {
			rst = append(rst, l)
			rst = append(rst, tmp)
		}
	}
	return rst
}