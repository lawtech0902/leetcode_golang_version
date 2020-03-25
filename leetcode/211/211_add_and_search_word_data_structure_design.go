/*
__author__ = 'lawtech'
__date__ = '2020/03/25 1:46 下午'
*/

package _211

// Trie树 + '.'特殊处理
type WordDictionary struct {
	next   map[rune]*WordDictionary
	isWord bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		next:   make(map[rune]*WordDictionary),
		isWord: false,
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	for _, v := range word {
		if this.next[v] == nil {
			this.next[v] = &WordDictionary{
				next:   make(map[rune]*WordDictionary),
				isWord: false,
			}
		}

		this = this.next[v]
	}

	this.isWord = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	for k, v := range word {
		if v == '.' {
			for _, v := range this.next {
				if v.Search(word[k+1:]) {
					return true
				}
			}

			return false
		} else {
			if this.next[v] == nil {
				return false
			}

			this = this.next[v]
		}
	}

	return this.isWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
