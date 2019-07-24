package main

func main() {

}

func getRoot(arr []int8, ind int8) int8 {
	affectInd := make([]int8, 0, 3)
	root := ind
	affectInd = append(affectInd, ind)
	for arr[ind] != ind {
		root = arr[ind]
		ind = root
		affectInd = append(affectInd, root)
	}
	return root
}

func equationsPossible(equations []string) bool {
	rst := make([]int8, 26, 26)
	for ind := range rst {
		rst[ind] = int8(ind)
	}
	for _, e := range equations {
		if e[1:3] == "==" {
			left := int8(e[0]) - 'a'
			right := int8(e[3]) - 'a'
			leftRoot := getRoot(rst, left)
			rightRoot := getRoot(rst, right)
			if leftRoot != rightRoot {
				if leftRoot < rightRoot {
					rst[rightRoot] = leftRoot
				} else {
					rst[leftRoot] = rightRoot
				}
			}
		}
	}
	for _, e := range equations {
		if e[1:3] == "!=" {
			left := int8(e[0]) - 'a'
			right := int8(e[3]) - 'a'
			leftRoot := getRoot(rst, left)
			rightRoot := getRoot(rst, right)
			if leftRoot == rightRoot {
				return false
			}
		}
	}
	return true
}

