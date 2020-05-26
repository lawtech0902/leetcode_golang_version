/*
__author__ = 'lawtech'
__date__ = '2020/05/26 2:07 下午'
*/

package _468

import (
	"fmt"
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if isValidIPv4(IP) {
		return "IPv4"
	}

	if isValidIPv6(IP) {
		return "IPv6"
	}

	return "Neither"
}

func isValidIPv4(IP string) bool {
	strArr := strings.Split(IP, ".")
	if len(strArr) != 4 {
		return false
	}

	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil || num > 255 || num < 0 {
			return false
		}

		newStr := fmt.Sprint(num)
		if str != newStr {
			return false
		}
	}

	return true
}

func isValidIPv6(IP string) bool {
	strArr := strings.Split(IP, ":")
	if len(strArr) != 8 {
		return false
	}

	for _, str := range strArr {
		if len(str) == 0 || len(str) > 4 {
			return false
		}

		for i := 0; i < len(str); i++ {
			if !(str[i] >= '0' && str[i] <= '9') &&
				!(str[i] >= 'a' && str[i] <= 'f') &&
				!(str[i] >= 'A' && str[i] <= 'F') {
				return false
			}
		}
	}

	return true
}
