/*
__author__ = 'lawtech'
__date__ = '2020/03/05 4:45 下午'
*/

package _134

func canCompleteCircuit(gas []int, cost []int) int {
	totalTank, curTank := 0, 0
	startStation := 0

	for i := range gas {
		totalTank += gas[i] - cost[i]
		curTank += gas[i] - cost[i]

		// 如果不能到达
		if curTank < 0 {
			// 选择下一个加油站作为始发站
			startStation = i + 1
			// 清空油箱
			curTank = 0
		}
	}

	if totalTank >= 0 {
		return startStation
	} else {
		return -1
	}
}
