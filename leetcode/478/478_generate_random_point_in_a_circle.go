/*
__author__ = 'lawtech'
__date__ = '2020/05/28 10:46 上午'
*/

package _478

import (
	"math"
	"math/rand"
)

type Solution struct {
	radius  float64
	xCenter float64
	yCenter float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		radius:  radius,
		xCenter: x_center,
		yCenter: y_center,
	}
}

func (this *Solution) RandPoint() []float64 {
	x0, y0 := this.xCenter-this.radius, this.yCenter-this.radius
	for {
		xg := x0 + rand.Float64()*this.radius*2
		yg := y0 + rand.Float64()*this.radius*2
		if math.Sqrt(math.Pow(xg-this.xCenter, 2)+math.Pow(yg-this.yCenter, 2)) <= this.radius {
			return []float64{xg, yg}
		}
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
