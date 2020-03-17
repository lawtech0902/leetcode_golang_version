/*
__author__ = 'lawtech'
__date__ = '2020/03/17 10:19 上午'
*/

package _173

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	next   *BSTIterator
	val    int
	length int
}

func Constructor(root *TreeNode) BSTIterator {
	// 中序遍历转为升序链表
	head := &BSTIterator{}

	var (
		dfs func(root *TreeNode)
	)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Right)
		head.length++
		cur := &BSTIterator{nil, root.Val, 0}
		cur.next = head.next
		head.next = cur
		dfs(root.Left)
	}

	dfs(root)
	return *head
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if this.next != nil {
		v := this.next.val
		this.next = this.next.next
		this.length--
		return v
	}

	return 0
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.length > 0 {
		return true
	}

	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
