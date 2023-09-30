package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) ToSlice() []int{
	s := make([]int, 0, 50)

	if l == nil {
		return s
	}

	s = append(s, l.Val)
	for l.Next != nil {
		s = append(s, l.Next.Val)
		l = l.Next
	}
	return s
}

func (l *ListNode) Dup() *ListNode {
	head := &ListNode{}
	head.Val = l.Val
	if l.Next != nil {
		head.Next = l.Next.Dup()
	}
	return head
}

