/*
__author__ = 'lawtech'
__date__ = '2020/04/27 5:01 下午'
*/

package _336

type node struct {
	child              [26]*node
	idx                int
	palindromeBranches []int
}

func insert(root *node, word string, idx int) {
	cur := root
	for i := len(word) - 1; i >= 0; i-- {
		c := int(word[i] - 'a')
		n := cur.child[c]
		if n == nil {
			n = &node{idx: -1}
			cur.child[c] = n
		}
		cur = n
		if i == 0 {
			cur.idx = idx
		}
	}
}

func buildTrie(words []string) *node {
	root := &node{idx: -1}
	for i, word := range words {
		insert(root, word, i)
	}
	return root
}

func palindrome(word string) bool {
	left := 0
	right := len(word) - 1
	for left <= right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func branches(prefix string, n *node) ([]string, []int) {
	var strs []string
	var ints []int
	if n.idx >= 0 {
		strs = append(strs, prefix)
		ints = append(ints, n.idx)
	}

	for i, child := range n.child {
		if child != nil {
			r0, r1 := branches(prefix+string([]byte{byte(i) + 'a'}), child)
			strs = append(strs, r0...)
			ints = append(ints, r1...)
		}
	}
	return strs, ints
}

func palindromeBranches(n *node) []int {
	if n.palindromeBranches != nil {
		return n.palindromeBranches
	}

	var ret []int
	strs, ints := branches("", n)
	for i := range strs {
		if palindrome(strs[i]) {
			ret = append(ret, ints[i])
		}
	}

	if ret == nil {
		n.palindromeBranches = make([]int, 0)
	} else {
		n.palindromeBranches = ret
	}
	return ret
}

func palindromePairs(words []string) [][]int {
	root := buildTrie(words)
	var ret [][]int
	for idx, word := range words {
		cur := root
		lastIdx := cur.idx
		for len(word) > 0 {
			if cur.idx >= 0 {
				if palindrome(word) {
					if lastIdx != idx {
						ret = append(ret, []int{idx, cur.idx})
					}
				}
			}

			cur = cur.child[int(word[0]-'a')]
			if cur == nil {
				break
			}
			lastIdx = cur.idx
			word = word[1:]
		}

		if cur != nil {
			// traverse the rest of the trie,
			// a match if any of the branches is a palindrome
			branches := palindromeBranches(cur)
			for _, b := range branches {
				if idx != b {
					ret = append(ret, []int{idx, b})
					if words[idx] == "" {
						// handle special case for the "" word
						ret = append(ret, []int{b, idx})
					}
				}
			}
		}
	}
	return ret
}
