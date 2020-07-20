/*
__author__ = 'lawtech'
__date__ = '2020/07/17 2:20 下午'
*/

package _729

// 朴素BST
type Node struct {
	left, right *Node
	start, end  int
}

func (n *Node) insert(node *Node) bool {
	if node.start >= n.end {
		if n.right == nil {
			n.right = node
			return true
		}

		return n.right.insert(node)
	}

	if node.end <= n.start {
		if n.left == nil {
			n.left = node
			return true
		}

		return n.left.insert(node)
	}

	return false
}

type MyCalendar struct {
	root *Node
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	node := &Node{
		start: start,
		end:   end,
	}

	if this.root == nil {
		this.root = node
		return true
	}

	return this.root.insert(node)
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

// 暴力
type MyCalendar1 struct {
	calendar [][]int
}

func Constructor1() MyCalendar {
	return MyCalendar{calendar: make([][]int, 0)}
}

func (this *MyCalendar1) Book1(start int, end int) bool {
	for _, c := range this.calendar {
		if c[0] < end && start < c[1] {
			return false
		}
	}

	this.calendar = append(this.calendar, []int{start, end})
	return true
}
