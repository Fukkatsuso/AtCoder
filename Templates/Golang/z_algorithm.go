package main

// 文字列sの先頭からの部分文字列と一致する部分文字列の開始地点に，
// その最大一致文字列長が格納された配列aを返す
// O(n)
func ZAlgorithm(s string) []int {
	n := len(s)
	a := make([]int, n)
	a[0] = n
	for i, j := 1, 0; i < n; {
		for i+j < n && s[j] == s[i+j] {
			j++
		}
		a[i] = j
		if j == 0 {
			i++
			continue
		}
		k := 1
		for i+k < n && k+a[k] < j {
			a[i+k] = a[k]
			k++
		}
		i += k
		j -= k
	}
	return a
}
