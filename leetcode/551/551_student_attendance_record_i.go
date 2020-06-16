/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-16 10:09:28
 */

package _551

import "strings"

func checkRecord(s string) bool {
	return !(strings.Count(s, "A") > 1) && !(strings.Contains(s, "LLL"))
}
