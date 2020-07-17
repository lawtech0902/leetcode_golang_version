/*
__author__ = 'lawtech'
__date__ = '2020/07/16 4:06 下午'
*/

package _722

func removeComments(source []string) []string {
	var (
		res     []string
		inBlock bool
		curLine string
	)

	for _, line := range source {
		// 不处于注释块中
		if !inBlock {
			curLine = ""
		}

		lv := len(line)
		for i := 0; i < lv; {
			// 处于注释块中
			if inBlock {
				// 注释块结束
				if i+2 <= lv && line[i:i+2] == "*/" {
					inBlock = false
					i += 2
					continue
				}

				// 其他情况，直接忽略
				i++
				continue
			} else {
				// 注释块
				if i+2 <= lv && line[i:i+2] == "/*" {
					inBlock = true
					i += 2
					continue
				}

				// 行注释
				if i+2 <= lv && line[i:i+2] == "//" {
					// 一整行，后面的就不看了
					break
				}

				// 其他情况，添加
				curLine += string(line[i])
				i++
			}
		}

		if !inBlock && curLine != "" {
			res = append(res, curLine)
		}
	}

	return res
}
