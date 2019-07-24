package main

func main()  {

}


type TreeNode struct {
    Val int
	Left *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sl, pl := doMaxPathSum(root.Right)
	sr, pr := doMaxPathSum(root.Left)
	noneg(&sl)
	noneg(&sr)
	noneg(&pl)
	noneg(&pr)
	mp := intMax(pl, pr)
	ms := intMax(sl, sr)
	mp = intMax(sl + sr + root.Val, mp)
	rst := intMax(mp, intMax(ms, ms + root.Val))
	if rst == 0 {
		return maxTree(root)
	} else {
		return rst
	}
}

func maxTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left != nil {
		return intMax(maxTree(root.Left), root.Val)
	}
	if root.Left == nil && root.Right != nil {
		return intMax(maxTree(root.Right), root.Val)
	}
	if root.Left != nil && root.Right != nil {
		return intMax(intMax(maxTree(root.Left), maxTree(root.Right)), root.Val)
	}
	return root.Val
}

func noneg(i *int)  {
	if *i < 0 {
		*i = 0
	}
}

func intMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func doMaxPathSum(root *TreeNode) (single int, path int) {
	if root == nil {
		return 0, 0
	}
	sl, pl := doMaxPathSum(root.Right)
	sr, pr := doMaxPathSum(root.Left)
	noneg(&sl)
	noneg(&sr)
	noneg(&pl)
	noneg(&pr)
	mp := intMax(pl, pr)
	ms := intMax(sl, sr)
	mp = intMax(sl + sr + root.Val, mp)
	return ms + root.Val, mp
}