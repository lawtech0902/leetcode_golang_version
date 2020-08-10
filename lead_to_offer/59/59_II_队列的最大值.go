/*
__author__ = 'lawtech'
__date__ = '2020/08/10 11:00 上午'
*/

package _59

// 有点类似最小栈
type MaxQueue struct {
	data []int
	max  []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		data: make([]int, 0),
		max:  make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}

	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.data = append(this.data, value)
	for len(this.max) != 0 && value > this.max[len(this.max)-1] {
		this.max = this.max[:len(this.max)-1]
	}

	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	n := -1
	if len(this.data) != 0 {
		n = this.data[0]
		this.data = this.data[1:]
		if this.max[0] == n {
			this.max = this.max[1:]
		}
	}

	return n
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
