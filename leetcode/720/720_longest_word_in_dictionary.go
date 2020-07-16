/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-16 11:23:56
 */

package _720

import (
	"sort"
	"strings"
)

// 排序 + map
func longestWord(words []string) string {
	sort.Strings(words)
	var res string
	set := make(map[string]struct{}, len(words))
	for _, word := range words {
		if _, ok := set[word[:len(word)-1]]; !ok && len(word) != 1 {
			continue
		}

		set[word] = struct{}{}
		if len(word) > len(res) {
			res = word
		}
	}

	return res
}

// 前缀树 + dfs
func longestWord1(words []string) string {
	trie := Trie{}
	trie.InsertWords(words)
	return trie.dfs(words)
}

type Trie struct {
	ending int
	next   [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string, i int) {
	tmp := this
	for _, v := range word {
		n := v - 'a'
		if tmp.next[n] == nil {
			tmp.next[n] = &Trie{}
		}

		tmp = tmp.next[n]
	}

	tmp.ending = i
}

func (this *Trie) InsertWords(words []string) {
	for k, v := range words {
		this.Insert(v, k+1)
	}
}

func (this *Trie) dfs(words []string) string {
	res := ""
	stack := []*Trie{this}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.ending > 0 || node == this {
			if node != this {
				tmpWord := words[node.ending-1]
				if len(tmpWord) > len(res) || (len(tmpWord) == len(res) && strings.Compare(res, tmpWord) == 1) {
					res = tmpWord
				}
			}

			for _, v := range node.next {
				if v != nil {
					stack = append(stack, v)
				}
			}
		}
	}

	return res
}
