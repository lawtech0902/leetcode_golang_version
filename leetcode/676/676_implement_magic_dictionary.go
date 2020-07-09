/*
__author__ = 'lawtech'
__date__ = '2020/07/09 1:49 下午'
*/

package _676

// 暴力
type MagicDictionary struct {
	hash map[int][]string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{hash: make(map[int][]string)}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for _, v := range dict {
		this.hash[len(v)] = append(this.hash[len(v)], v)
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	words := this.hash[len(word)]
	for _, v := range words {
		t := 0
		for i := 0; i < len(word); i++ {
			if v[i] != word[i] {
				t++
			}

			if t > 1 {
				break
			}
		}

		if t == 1 {
			return true
		}
	}

	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */
