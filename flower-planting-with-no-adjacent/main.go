package main

import "fmt"

func main()  {
	a := 5
	b := [][]int{[]int{4,1}, []int{4,2}, []int{4,3}, []int{2,5}, []int{1,2}, []int{1, 5}}
	//a := 4
	//b := [][]int{[]int{1,2}, []int{2,3}, []int{3,4}, []int{4,1}, []int{1,3}, []int{2,4}}
	fmt.Println(gardenNoAdj(a, b))
}

func getRemain(l, r int, ls, rs map[int]bool) (int, int) {
	tmp := map[int]bool{1: true, 2: true, 3: true, 4: true}
	if r == 0 {
		if l == 0 {
			for t := range tmp {
				if !ls[t] {
					l = t
					break
				}
			}
		}
		delete(tmp, l)
		for t := range tmp {
			if !rs[t] {
				r = t
				break
			}
		}
	} else {
		if l == 0 {
			delete(tmp, r)
			for t := range tmp {
				if !ls[t] {
					l = t
					break
				}
			}
		}

	}


	return l, r
}

func trans(ss map[int]bool, rr []int) map[int]bool {
	rst := make(map[int]bool, len(ss))
	for s := range ss {
		if rr[s] != 0 {
			rst[rr[s]] = true
		}
	}
	return rst
}

func gardenNoAdj(N int, paths [][]int) []int {
	mm := make(map[int]map[int]bool, N)
	for _, edge := range paths {
		l := edge[0] - 1
		r := edge[1] - 1
		if _, exist := mm[l]; !exist {
			mm[l] = make(map[int]bool, 4)
		}
		if _, exist := mm[r]; !exist {
			mm[r] = make(map[int]bool, 4)
		}
		mm[l][r] = true
		mm[r][l] = true
	}
	rst := make([]int, N)
	for _, edge := range paths {
		l := edge[0] - 1
		r := edge[1] - 1
		if rst[l] != 0 && rst[r] != 0 {
			continue
		}
		ls := mm[l]
		rs := mm[r]
		ll, rr := getRemain(rst[l], rst[r], trans(ls, rst), trans(rs, rst))
		rst[l] = ll
		rst[r] = rr
	}
	for ind := range rst {
		if rst[ind] == 0 {
			rst[ind]++
		}
	}
	return rst
}