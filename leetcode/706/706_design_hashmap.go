/*
__author__ = 'lawtech'
__date__ = '2020/07/15 11:12 上午'
*/

package _706

type MyHashMap struct {
	capacity int
	data     []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	capacity := 1000001
	data := make([]int, capacity)
	for i := range data {
		data[i] = -1
	}

	return MyHashMap{
		capacity: capacity,
		data:     data,
	}
}

func (this *MyHashMap) Hash(key int) int {
	return key % this.capacity
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.data[this.Hash(key)] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.data[this.Hash(key)]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.data[this.Hash(key)] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
