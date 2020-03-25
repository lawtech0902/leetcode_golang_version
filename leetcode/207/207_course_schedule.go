/*
__author__ = 'lawtech'
__date__ = '2020/03/25 9:53 上午'
*/

package _207

/*
针对所有的V建立一个入度表。
并构造一个queue，当某个V的入度为0时，加入queue中。
当queue长度不为0时，从queue取出一个节点，并将该节点指向的节点的入度-1。若-1后，某个V满足入度==0，则将其放入queue中。
每次有元素出队列时，numCourse个数减一。
理论上，当queue为空时，应该刚好numCourse减到0.若不是，则说明有课程无法选修到。
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	list := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	queue := make([]int, 0)

	for i, _ := range prerequisites {
		src, target := prerequisites[i][0], prerequisites[i][1]
		inDegree[target] += 1
		if nil == list[src] {
			list[src] = []int{}
		}

		list[src] = append(list[src], target)
	}

	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		c := queue[0]
		queue = queue[1:]
		for _, target := range list[c] {
			inDegree[target] -= 1
			if inDegree[target] == 0 {
				queue = append(queue, target)
			}
		}

		numCourses--
	}

	return numCourses == 0
}
