// 解説AC
// イベントソートまでは思いついた

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

type Event struct {
	day, c int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, c := getInt(), getInt()

	events := make([]Event, 2*n)
	for i := 0; i < n; i++ {
		ai, bi, ci := getInt(), getInt(), getInt()
		events[2*i] = Event{ai, ci}
		events[2*i+1] = Event{bi + 1, -ci}
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].day < events[j].day
	})

	ans := 0
	cost := 0
	lastDay := 0
	for _, event := range events {
		today := event.day
		if today != lastDay {
			ans += min(c, cost) * (today - lastDay)
			lastDay = today
		}
		cost += event.c
	}

	puts(ans)
}
