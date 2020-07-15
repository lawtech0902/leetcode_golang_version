/*
__author__ = 'lawtech'
__date__ = '2020/07/15 10:33 上午'
*/

package _705

const (
	BucketSize = 731
	Deleted    = -1
)

type MyHashSet struct {
	bucket [BucketSize][]int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{}
}

func hash(key int) int {
	return key % BucketSize
}

func (this *MyHashSet) Add(key int) {
	index := hash(key)
	for _, v := range this.bucket[index] {
		if v == key {
			return
		}
	}
	this.bucket[index] = append(this.bucket[index], key)
}

func (this *MyHashSet) Remove(key int) {
	index := hash(key)
	for i, v := range this.bucket[index] {
		if v == key {
			this.bucket[index][i] = Deleted
			return
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	index := hash(key)
	for _, v := range this.bucket[index] {
		if v == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
