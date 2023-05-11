package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n, summ, elem int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &n)

	var eqPref map[int][]int = make(map[int][]int)
	eqPref[0] = append(eqPref[0], 0)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &elem)
		summ += elem
		eqPref[summ] = append(eqPref[summ], i)
		// fmt.Println(i, summ)
	}

	var segments [][2]int = make([][2]int, n)

	for _, elem := range eqPref {
		for i := 1; i < len(elem); i++ {
			var tmp [2]int
			tmp[0] = elem[i-1] + 1
			tmp[1] = elem[i]
			segments = append(segments, tmp)
		}
	}

	sort.Slice(segments, func(a, b int) bool {
		if segments[a][0] == segments[b][0] {
			return segments[a][1] < segments[b][1]
		}
		return segments[a][0] < segments[b][0]
	})

	var startIdx int = -1
	for i := 0; i < len(segments); i++ {
		if segments[i][0] != 0 {
			startIdx = i
			break
		}
	}

	var res int
	if startIdx != -1 {
		res = (segments[startIdx][0] - 0) * (n - segments[startIdx][1] + 1)
	}

	for i := startIdx + 1; i < len(segments); i++ {
		res += (segments[i][0] - segments[i-1][0]) * (n - segments[i][1] + 1)
	}

	// fmt.Println(eqPref)
	// fmt.Print(segments)
	fmt.Print(res)
}
