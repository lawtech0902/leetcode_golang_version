/*
__author__ = 'lawtech'
__date__ = '2020/1/13 5:00 下午'
*/

package _43

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// string转换成[]byte，容易取得相应位上的具体值
	bsNum1 := []byte(num1)
	bsNum2 := []byte(num2)

	temp := make([]int, len(num1)+len(num2))
	for i := 0; i < len(bsNum1); i++ {
		for j := 0; j < len(bsNum2); j++ {
			// 每个位上的乘积，可以直接存入 temp 中相应的位置
			temp[i+j+1] += int(bsNum1[i]-'0') * int(bsNum2[j]-'0')
		}
	}

	// 统一处理进位
	for i := len(temp) - 1; i > 0; i-- {
		temp[i-1] += temp[i] / 10
		temp[i] = temp[i] % 10
	}

	// num1 和 num2 较小的时候，temp的首位为0
	// 为避免输出结果以0开头，需要去掉temp的0首位
	if temp[0] == 0 {
		temp = temp[1:]
	}

	// 转换结果
	// temp 选用为[]int，而不是[]byte，是因为
	// Go中，byte的基础结构是uint8，最大值为255。
	// 不考虑进位的话，temp会溢出
	res := make([]byte, len(temp))
	for i := 0; i < len(temp); i++ {
		res[i] = '0' + byte(temp[i])
	}

	return string(res)
}
