package main

import (
	"fmt"
	"strconv"
)

func main()  {
	//a := [][]int{[]int{1,1}, []int{3,2}, []int{5,3}, []int{4,1}, []int{2,3}, []int{1,4}}
	a := [][]int{[]int{0, 0}, []int{94911151,94911150},[]int{94911152,94911151}}
	fmt.Println(maxPoints(a))

}

func gcd(a, b int) int {
	if a % b == 0 {
		return b
	}
	return gcd(b, a %b)
}

func calRate(p1, p2 []int) string {
	if p1[0] == p2[0] {
		return "inf" + strconv.Itoa(p1[0])
	} else {
		x := p2[0] - p1[0]
		y := p2[1] - p1[1]
		if x < 0 {
			x = -x
			y = -y
		}
		slopGcd := gcd(y, x)
		x = x / slopGcd
		y = y / slopGcd
		slop := float64(y) / float64(x)
		xc := float64(p1[0]) - float64(p1[1]) / slop
		return strconv.Itoa(x) + strconv.Itoa(y) + strconv.FormatFloat(xc, 'f', 15, 64)
	}
}

func maxPoints(points [][]int) int {
	l := len(points)
	if l == 0 {
		return 0
	}
	r := make(map[string]map[int]bool, 10)
	for index := 0; index < l; index++ {
		for ind := index + 1; ind < l; ind++ {
			slop := calRate(points[index], points[ind])
			if _, exist := r[slop]; !exist {
				r[slop] = make(map[int]bool, 2)
			}
			r[slop][index] = true
			r[slop][ind] = true
		}
	}
	max := 0
	for _, num := range r {
		if len(num) > max {
			max = len(num)
		}
	}
	if max == 0 {
		max = 1
	}
	return max
}