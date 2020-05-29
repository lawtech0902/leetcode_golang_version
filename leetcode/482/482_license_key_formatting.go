/*
__author__ = 'lawtech'
__date__ = '2020/05/29 11:07 上午'
*/

package _482

import "strings"

func licenseKeyFormatting(S string, K int) string {
	S = strings.ToUpper(S)
	S = strings.Replace(S, "-", "", -1)

	for i := len(S) - 1 - K; i >= 0; i = i - K {
		S = S[0:i+1] + "-" + S[i+1:]
	}

	return S
}
