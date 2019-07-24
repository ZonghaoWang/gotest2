package main

import "fmt"

func main() {
	//a := [][]int{[]int{2, 1},[]int{3, 1},[]int{4, 2}, []int{1, 4}}
	a := [][]int{[]int{4, 1}, []int{1, 5}, []int{4, 2}, []int{5, 1}, []int{4,3}}
	fmt.Println(findRedundantDirectedConnection(a))
}



func findRedundantDirectedConnection(edges [][]int) []int {
	in := make([]int, len(edges))
	for _, edge := range edges {
		in[edge[1]-1]++
	}
	root := -1
	lp := -1
	for ind, inDegree := range in {
		if inDegree == 0 {
			root = ind
		}
		if inDegree == 2 {
			lp = ind
		}
	}
	if root == -1 {
		// 无根节点这种情况,之前有出度的顶点最后一次入度
		mm := make(map[int]bool, len(edges))
		for _, edge := range edges {
			if mm[edge[1] - 1] {
				return edge
			}
			mm[edge[0] - 1] = true
		}

	}
	// 能不能从lp追溯到root
	arr := make([]int, len(edges))
	for ind := range arr {
		arr[ind] = ind
	}
	var first, second []int
	for _, edge := range edges {
		if edge[1] - 1 == lp && arr[edge[1] - 1] == edge[1] - 1 {
			// 第一次的入度
			first = edge
		}
		if edge[1] - 1 == lp && arr[edge[1] - 1] != edge[1] - 1 {
			second = edge
			continue
		}
		arr[edge[1] - 1] = edge[0] - 1
	}
	// 能不能从lp找打root
	tmp := lp
	for arr[tmp] != root {
		tmp = arr[tmp]
		if tmp == lp {
			return first
		}
	}
	return second
}