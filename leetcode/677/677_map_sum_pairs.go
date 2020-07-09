/*
__author__ = 'lawtech'
__date__ = '2020/07/09 2:04 下午'
*/

package _677

// 前缀树
type TrieNode struct {
	children map[byte]*TrieNode
	sum      int
}

type MapSum struct {
	hash map[string]int
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		hash: make(map[string]int),
		root: &TrieNode{children: make(map[byte]*TrieNode)},
	}
}

func (this *MapSum) Insert(key string, val int) {
	diff := val - this.hash[key]
	this.hash[key] = val
	cur := this.root
	for i := range key {
		if _, ok := cur.children[key[i]]; !ok {
			trie := TrieNode{children: make(map[byte]*TrieNode)}
			cur.children[key[i]] = &trie
		}

		cur = cur.children[key[i]]
		cur.sum += diff
	}
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.root
	for i := range prefix {
		if _, ok := cur.children[prefix[i]]; !ok {
			return 0
		}

		cur = cur.children[prefix[i]]
	}

	return cur.sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
