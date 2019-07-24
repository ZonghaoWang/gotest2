package main

import "fmt"

func main() {
	a := [][]int{[]int{1,2},[]int{1,3},[]int{2,3}}
	fmt.Println(findRedundantConnection(a))
}


func findRoot(arr []int, ind int) int {
	affertArr := make([]int, 0, 10)
	affertArr = append(affertArr, ind)
	var parent = arr[ind]
	var root = parent
	for parent != ind {
		affertArr = append(affertArr, parent)
		ind = parent
		parent = arr[ind]
		root = parent
	}
	for _, v := range affertArr {
		arr[v] = root
	}
	return arr[ind]
}

func findRedundantConnection(edges [][]int) []int {
	arr := make([]int, len(edges))
	for ind := range arr {
		arr[ind] = ind
	}
	for _, edge := range edges {
		left := edge[0] - 1
		right := edge[1] - 1
		leftRoot := findRoot(arr, left)
		rightRoot := findRoot(arr, right)
		if leftRoot == rightRoot {
			return edge
		}
		if leftRoot < rightRoot {
			arr[rightRoot] = leftRoot
		} else {
			arr[leftRoot] = rightRoot
		}
	}
	return nil
}
