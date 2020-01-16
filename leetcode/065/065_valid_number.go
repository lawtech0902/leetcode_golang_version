/*
__author__ := 'lawtech'
__date__ := '2020/1/16 3:22 下午'
*/

package _65

func isNumber(s string) bool {
	// 确定有穷状态自动机DFA，什么鬼东西
	INVALID := 0
	SPACE := 1
	SIGN := 2
	DIGIT := 3
	DOT := 4
	EXPONENT := 5
	//  0invalid,1space,2sign,3digit,4dot,5exponent,6num_inputs
	transitionTable := [][]int{
		{-1, 0, 3, 1, 2, -1},    // 0 no input or just spaces
		{-1, 8, -1, 1, 4, 5},    // 1 input is digits
		{-1, -1, -1, 4, -1, -1}, // 2 no digits in front just Dot
		{-1, -1, -1, 1, 2, -1},  // 3 sign
		{-1, 8, -1, 4, -1, 5},   // 4 digits and dot in front
		{-1, -1, 6, 7, -1, -1},  // 5 input 'e' or 'E'
		{-1, -1, -1, 7, -1, -1}, // 6 after 'e' input sign
		{-1, 8, -1, 7, -1, -1},  // 7 after 'e' input digits
		{-1, 8, -1, -1, -1, -1},
	} // 8 after valid input input space
	state := 0
	i := 0

	for i < len(s) {
		inputType := INVALID
		switch s[i] {
		case ' ':
			inputType = SPACE
		case '-', '+':
			inputType = SIGN
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			inputType = DIGIT
		case '.':
			inputType = DOT
		case 'e', 'E':
			inputType = EXPONENT
		}

		state = transitionTable[state][inputType]
		if state == -1 {
			return false
		} else {
			i += 1
		}
	}

	return state == 1 || state == 4 || state == 7 || state == 8
}
