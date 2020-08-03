/*
__author__ = 'lawtech'
__date__ = '2020/08/03 10:40 上午'
*/

package _35

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// hash
func copyRandomList(head *Node) *Node {
	hash := make(map[*Node]*Node)
	cur := head
	// 复制节点值
	for cur != nil {
		hash[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		hash[cur].Next = hash[cur.Next]
		hash[cur].Random = hash[cur.Random]
		cur = cur.Next
	}

	return hash[head]
}
