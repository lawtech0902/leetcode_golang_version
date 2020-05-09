/*
__author__ = 'lawtech'
__date__ = '2020/05/09 11:13 上午'
*/

package _381

import "math/rand"

type RandomizedCollection struct {
	set    map[int]*Node
	data   []int
	length int
}

type Node struct {
	val  int
	next *Node
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		set:    make(map[int]*Node),
		data:   make([]int, 0, 100),
		length: 0,
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (rc *RandomizedCollection) Insert(val int) bool {
	res := false
	if _, ok := rc.set[val]; !ok {
		res = true
	}

	rc.data = append(rc.data, val)
	newNode := &Node{
		val:  rc.length,
		next: rc.set[val],
	}

	rc.set[val] = newNode
	rc.length++
	return res
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (rc *RandomizedCollection) Remove(val int) bool {
	if _, ok := rc.set[val]; ok {
		index := rc.set[val].val
		lastVal := rc.data[len(rc.data)-1]
		rc.data[index] = lastVal
		rc.set[lastVal].val = index
		if rc.set[lastVal].next != nil && rc.set[lastVal].next.val > rc.set[lastVal].val {
			rc.set[lastVal].next.val, rc.set[lastVal].val = rc.set[lastVal].val, rc.set[lastVal].next.val
		}

		rc.data = rc.data[:len(rc.data)-1]
		rc.set[val] = rc.set[val].next
		if rc.set[val] == nil {
			delete(rc.set, val)
		}

		rc.length--
		return true
	}

	return false
}

/** Get a random element from the collection. */
func (rc *RandomizedCollection) GetRandom() int {
	if rc.length > 0 {
		index := rand.Intn(rc.length)
		return rc.data[index]
	}

	return -1
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
