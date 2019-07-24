package main

import "fmt"

func main()  {
	a := []int{2,3,1,1,4}
	fmt.Println(canJump(a))
}

func intMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func canJump(nums []int) bool {
	l := len(nums)
	max := 0
	for ind := 0; ind < l - 1; ind++ {
		max = intMax(ind + nums[ind], max)
		if max <= ind {
			return false
		}
	}
	return true
}