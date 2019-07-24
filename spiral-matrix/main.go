package main

import "fmt"

func main()  {
	a := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}
	fmt.Println(spiralOrder(a))
}

type Direction int
const (
	_ Direction = iota
	START
	RIGHT
	DOWN
	LEFT
	TOP
)

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	l, t := 0, 0
	dire := START
	rst := make([]int, 0, m * n)
	r, b := n - 1, m - 1
	x, y := 0, 0
	status := false
	for l <= r && t <= b {
		switch dire {
		case START:
			rst = append(rst, matrix[0][0])
			if r == y {
				if x == b {
					return rst
				}
				x++
				dire = DOWN
			} else {
				y++
				dire = RIGHT
			}
			status = true
		case RIGHT:
			if status {
				if y == l {
					y++
					l++
					status = false
					continue
				}
			}
			rst = append(rst, matrix[x][y])
			if y == r {
				dire = DOWN
				status = true
				continue
			}
			y++
		case DOWN:
			if status {
				if x == t {
					x++
					t++
					status = false
					continue
				}
			}
			rst = append(rst, matrix[x][y])
			if x == b {
				dire = LEFT
				status = true
				continue
			}
			x++
		case LEFT:
			if status {
				if y == r {
					y--
					r--
					status = false
					continue
				}
			}
			rst = append(rst, matrix[x][y])
			if y == l {
				dire = TOP
				status = true
				continue
			}
			y--
		case TOP:
			if status {
				if x == b {
					x--
					b--
					status = false
					continue
				}
			}
			rst = append(rst, matrix[x][y])
			if x == t {
				status = true
				dire = RIGHT
				continue
			}
			x--
		}
	}
	return rst
}
