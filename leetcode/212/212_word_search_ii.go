/*
__author__ = 'lawtech'
__date__ = '2020/03/25 2:00 下午'
*/

package _212

import "fmt"

type TrieNode struct {
	val  rune
	word string
	next map[rune]*TrieNode
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{next: make(map[rune]*TrieNode)}
	for _, word := range words {
		p := root
		for i, b := range word {
			if _, ok := p.next[b]; !ok {
				p.next[b] = &TrieNode{
					val:  b,
					next: make(map[rune]*TrieNode),
				}
			}

			if i == len(word)-1 {
				p.next[b].word = word
			}

			p = p.next[b]
		}
	}

	var res []string

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, &res, root, i, j, map[string]bool{})
		}
	}

	return res
}

func dfs(board [][]byte, res *[]string, node *TrieNode, i, j int, visited map[string]bool) {
	k := fmt.Sprintf("%v_%v", i, j)
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || visited[k] {
		return
	}

	nextNode, ok := node.next[rune(board[i][j])]
	if !ok {
		return
	}

	if len(nextNode.word) > 0 {
		*res = append(*res, nextNode.word)
		nextNode.word = ""
	}

	visited[k] = true
	dfs(board, res, nextNode, i+1, j, visited) // down
	dfs(board, res, nextNode, i-1, j, visited) // up
	dfs(board, res, nextNode, i, j+1, visited) // right
	dfs(board, res, nextNode, i, j-1, visited) // left
	visited[k] = false
}
