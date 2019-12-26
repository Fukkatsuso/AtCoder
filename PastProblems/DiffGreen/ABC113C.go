// 答え見た

package main

import (
	"fmt"
	"sort"
)

// 入力順, 誕生年, 県番号, id
type City struct {
	num, year, pref int
	id              string
}

// 入力順でソート
type CitiesByNum []*City

func (cn CitiesByNum) Len() int {
	return len(cn)
}

func (cn CitiesByNum) Less(i, j int) bool {
	return cn[i].num < cn[j].num
}

func (cn CitiesByNum) Swap(i, j int) {
	cn[i], cn[j] = cn[j], cn[i]
}

// 誕生年でソート
type CitiesByYear []*City

func (cy CitiesByYear) Len() int {
	return len(cy)
}

func (cy CitiesByYear) Less(i, j int) bool {
	if cy[i].pref == cy[j].pref {
		return cy[i].year < cy[j].year
	}
	return cy[i].pref < cy[j].pref
}

func (cy CitiesByYear) Swap(i, j int) {
	cy[i], cy[j] = cy[j], cy[i]
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	cities := make([]*City, m)
	for i := range cities {
		var p, y int
		fmt.Scan(&p, &y)
		cities[i] = &City{i, y, p, ""}
	}
	sort.Sort(CitiesByYear(cities))

	curPref, curOrder := 0, 0
	for i := range cities {
		if cities[i].pref != curPref {
			curOrder = 1
			curPref = cities[i].pref
		} else {
			curOrder++
		}
		cities[i].id = fmt.Sprintf("%06d%06d", cities[i].pref, curOrder)
	}

	sort.Sort(CitiesByNum(cities))

	for i := range cities {
		fmt.Println(cities[i].id)
	}
}
