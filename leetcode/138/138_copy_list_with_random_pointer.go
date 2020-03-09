/*
__author__ = 'lawtech'
__date__ = '2018/8/27 下午7:04'
*/

package _138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	ptr := head
	for ptr != nil {
		newNode := &Node{Val: ptr.Val}

		newNode.Next = ptr.Next
		ptr.Next = newNode
		ptr = newNode.Next
	}

	ptr = head

	for ptr != nil {
		if ptr.Random != nil {
			ptr.Next.Random = ptr.Random.Next
		} else {
			ptr.Next.Random = nil
		}

		ptr = ptr.Next.Next
	}

	ptrOldList := head
	ptrNewList := head.Next
	headOld := head.Next

	for ptrOldList != nil {
		ptrOldList.Next = ptrOldList.Next.Next
		if ptrNewList.Next != nil {
			ptrNewList.Next = ptrNewList.Next.Next
		} else {
			ptrNewList.Next = nil
		}

		ptrOldList = ptrOldList.Next
		ptrNewList = ptrNewList.Next
	}

	return headOld
}
