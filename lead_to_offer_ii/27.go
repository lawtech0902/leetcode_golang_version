package lead_to_offer_ii

// O(n) O(n)
func isPalindrome1(head *ListNode) bool {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}

	return true
}

// O(n) O(1)
func isPalindrome2(head *ListNode) bool {
	// 快慢指针找中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转后半段链表
	var pre *ListNode
	for slow != nil {
		temp := slow.Next
		slow.Next = pre
		pre = slow
		slow = temp
	}

	for pre != nil && head != nil {
		if pre.Val != head.Val {
			return false
		}

		pre = pre.Next
		head = head.Next
	}

	return true
}
