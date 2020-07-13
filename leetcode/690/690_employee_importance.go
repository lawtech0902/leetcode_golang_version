/*
__author__ = 'lawtech'
__date__ = '2020/07/13 11:14 上午'
*/

package _690

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	var (
		emap = make(map[int]*Employee)
		dfs  func(eid int) int
	)

	for _, e := range employees {
		emap[e.Id] = e
	}

	dfs = func(eid int) int {
		e := emap[eid]
		res := e.Importance
		for _, sid := range e.Subordinates {
			res += dfs(sid)
		}

		return res
	}

	return dfs(id)
}
