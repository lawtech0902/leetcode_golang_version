/*
__author__ = 'lawtech'
__date__ = '2020/03/25 10:37 上午'
*/

package _208

// map
type Trie struct {
	next   map[rune]*Trie
	isWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		next:   make(map[rune]*Trie),
		isWord: false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.next[v] == nil {
			this.next[v] = &Trie{
				next:   make(map[rune]*Trie),
				isWord: false,
			}
		}

		this = this.next[v]
	}

	this.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.prefixNode(word)
	return node != nil && node.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.prefixNode(prefix)
	return node != nil
}

func (this *Trie) prefixNode(prefix string) *Trie {
	for _, v := range prefix {
		if this.next[v] == nil {
			return nil
		}

		this = this.next[v]
	}

	return this
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
