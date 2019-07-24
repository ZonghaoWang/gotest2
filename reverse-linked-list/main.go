package main

func main()  {

}


type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	var newHead = &ListNode{}
	for p != nil {
		tmp := newHead.Next
		newHead.Next = p
		pn := p.Next
		p.Next = tmp
		p = pn
	}
	return newHead.Next
}