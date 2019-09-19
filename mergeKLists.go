package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	var curr *ListNode
	for {
		var min *ListNode
		var index int
		for i := range lists {
			if min == nil || (lists[i] != nil && lists[i].Val < min.Val) {
				min = lists[i]
				index = i
			}
		}
		if min == nil {
			break
		}
		if head == nil {
			head = min
			curr = head
		} else {
			curr.Next = min
			curr = curr.Next
		}
		lists[index] = lists[index].Next
	}

	return head
}
