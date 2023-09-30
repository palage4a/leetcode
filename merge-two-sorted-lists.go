package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{}
	if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l1 == nil && l2 == nil {
		return nil
	}

	if l1.Val <= l2.Val {
		h.Val = l1.Val
		if l1.Next != nil {
			h.Next = mergeTwoLists(l1.Next, l2)
		} else {
			h.Next = &ListNode{l2.Val, nil}
		}
	}

	if l1.Val > l2.Val {
		h.Val = l2.Val
		if l2.Next != nil {
			h.Next = mergeTwoLists(l1, l2.Next)
		} else {
			h.Next = &ListNode{l1.Val, nil}
		}
	}

	return h
}
