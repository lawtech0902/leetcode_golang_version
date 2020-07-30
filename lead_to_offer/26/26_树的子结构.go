/*
__author__ = 'lawtech'
__date__ = '2020/07/30 10:34 上午'
*/

package _26

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 注意：子结构和子树不一样，和572不一样
// 牛逼的代码
func isSubStructure(s, t *TreeNode) bool {
	return (s != nil && t != nil) && (recur(s, t) || isSubStructure(s.Left, t) || isSubStructure(s.Right, t))
}

func recur(s, t *TreeNode) bool {
	if t == nil {
		return true
	}

	if s == nil || s.Val != t.Val {
		return false
	}

	return recur(s.Left, t.Left) && recur(s.Right, t.Right)
}

// 菜鸡的代码
func isSubStructure1(s, t *TreeNode) bool {
	if s == nil || t == nil {
		return false
	}

	if s.Val == t.Val && helper(s.Left, t.Left) && helper(s.Right, t.Right) {
		return true
	}

	return isSubStructure1(s.Left, t) || isSubStructure1(s.Right, t)
}

func helper(s, t *TreeNode) bool {
	if t == nil {
		return true
	}

	if s == nil {
		return false
	}

	if s.Val == t.Val {
		return helper(s.Left, t.Left) && helper(s.Right, t.Right)
	} else {
		return false
	}
}
