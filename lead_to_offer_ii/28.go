package lead_to_offer_ii

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}

	dummy := &Node{
		Val:   0,
		Prev:  nil,
		Next:  root,
		Child: nil,
	}

	dfs(dummy, root)
	dummy.Next.Prev = nil
	return dummy.Next
}

func dfs(prev, cur *Node) *Node {
	if cur == nil {
		return prev
	}

	cur.Prev = prev
	prev.Next = cur
	tempNext := cur.Next
	tail := dfs(cur, cur.Child)
	cur.Child = nil
	return dfs(tail, tempNext)
}
