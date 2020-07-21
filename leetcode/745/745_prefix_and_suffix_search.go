/*
__author__ = 'lawtech'
__date__ = '2020/07/21 10:53 上午'
*/

package _745

type WordFilter struct {
	weight   int
	children [27]*WordFilter
}

func Constructor(words []string) WordFilter {
	wf := WordFilter{
		weight:   0,
		children: [27]*WordFilter{},
	}

	for weight, word := range words {
		key := word + "{" + word
		for i := 0; i <= len(word); i++ {
			wf.Add(key[i:], weight)
		}
	}

	return wf
}

func (this *WordFilter) Add(word string, weight int) {
	root := this
	for i := 0; i < len(word); i++ {
		if root.children[word[i]-'a'] == nil {
			root.children[word[i]-'a'] = &WordFilter{
				weight:   weight,
				children: [27]*WordFilter{},
			}
		}

		root = root.children[word[i]-'a']
		// 更新权重，因为可能出现重复的word，但是权重已经提升了
		root.weight = weight
	}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	root := this
	key := suffix + "{" + prefix
	for _, b := range key {
		if root.children[b-'a'] == nil {
			return -1
		}

		root = root.children[b-'a']
	}

	return root.weight
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
