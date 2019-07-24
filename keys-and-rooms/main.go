package main

import "fmt"

func main()  {
	a := [][]int{[]int{1}, []int{2}, []int{3}, []int{}}
	//a := [][]int{[]int{}}
	fmt.Println(canVisitAllRooms(a))
}


func canVisitAllRooms(rooms [][]int) bool {
	checked := make(map[int]bool, len(rooms))
	checked[0] = true
	for ind := 1; ind < len(rooms); ind++ {
		checked[ind] = false
	}
	if len(rooms) <= 1 {
		return true
	}
	for len(checked) > 0 {
		cks := make([]int, 0, 10)
		for ind, stat := range checked {
			if stat {
				cks = append(cks, ind)
			}
		}
		for _, indd := range cks {
			for _, ind := range rooms[indd] {
				if _, exist := checked[ind]; exist {
					checked[ind] = true
				}
			}
		}

		if len(cks) == 0 {
			return false
		}

		for _, ind := range cks {
			delete(checked, ind)
		}
	}
	return true
}