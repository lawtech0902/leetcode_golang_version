/*
__author__ = 'lawtech'
__date__ = '2020/05/09 10:41 上午'
*/

package _380

import "math/rand"

type RandomizedSet struct {
	buff  []int
	cache map[int]int // num -> idx
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		buff:  make([]int, 0),
		cache: make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.cache[val]; ok {
		return false
	}

	rs.buff = append(rs.buff, val)
	rs.cache[val] = len(rs.buff) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (rs *RandomizedSet) Remove(val int) bool {
	if idx, ok := rs.cache[val]; ok {
		rs.cache[rs.buff[len(rs.buff)-1]] = idx
		rs.buff[idx], rs.buff[len(rs.buff)-1] = rs.buff[len(rs.buff)-1], rs.buff[idx]
		rs.buff = rs.buff[:len(rs.buff)-1]
		delete(rs.cache, val)
		return true
	}

	return false
}

/** Get a random element from the set. */
func (rs *RandomizedSet) GetRandom() int {
	if len(rs.buff) > 0 {
		return rs.buff[rand.Intn(len(rs.buff))]
	}

	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
