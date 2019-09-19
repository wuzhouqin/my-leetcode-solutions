package main

func swapPairs(head *ListNode) *ListNode {
	dummy := ListNode{0, head}
	curr := &dummy
	for {
		if curr == nil {
			break
		}
		first := curr.Next
		if first == nil {
			break
		}
		second := first.Next
		if second == nil {
			break
		}
		first.Next = second.Next
		second.Next = first
		curr.Next = second
		curr = first
	}

	return dummy.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := ListNode{0, head}
	curr := &dummy
	ps := make([]*ListNode, k)
	for {
		if curr == nil || curr.Next == nil {
			break
		}
		ps[0] = curr.Next
		done := false
		for i := 1; i < k; i++ {
			ps[i] = ps[i-1].Next
			if ps[i] == nil {
				done = true
				break
			}
		}

		if done {
			break
		}

		curr.Next = ps[k-1]
		ps[0].Next = ps[k-1].Next
		for i := k - 1; i > 0; i-- {
			ps[i].Next = ps[i-1]
		}
		curr = ps[0]
	}
	return dummy.Next
}

func removeDuplicates(nums []int) int {
	count := 0
	p := 0
	for i, n := range nums {
		if i == 0 || n != nums[p-1] {
			nums[p] = n
			p++
			count++
		}
	}
	return count
}
