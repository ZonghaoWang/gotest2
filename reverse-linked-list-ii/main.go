package main

import "fmt"

func main()  {
	a1 := &ListNode{Val:1}
	a2 := &ListNode{Val:2}
	a3 := &ListNode{Val:3}
	a4 := &ListNode{Val:4}
	a5 := &ListNode{Val:5}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	fmt.Println(reverseBetween(a4, 1, 2))
}

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	cntM := 0
	cntN := 0
	begin := &ListNode{}
	begin.Next = head
	p := begin
	for p != nil {
		cntM ++
		cntN ++
		if cntM <= m {
			begin = p
			p = p.Next
			continue
		}
		if cntN > n {
			break
		}
		pn := p.Next
		if pn == nil {
			break
		}
		pnn := pn.Next
		p.Next = pnn
		pn.Next = begin.Next
		begin.Next = pn
	}

	if m == 1 {
		return begin.Next
	}
	return head
}