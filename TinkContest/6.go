// ОНА НЕ РАБОТАЕТ!!!

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n, maxSumm int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &n, &maxSumm)
	var data [][2]int = make([][2]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i][0], &data[i][1])
	}

	sort.Slice(data, func(a, b int) bool {
		if data[a][0] == data[b][0] {
			return data[a][1] < data[b][1]
		}
		return data[a][0] < data[b][0]
	})

	// fmt.Print(data)

	var minSumm int
	for i := 0; i < n/2; i++ {
		minSumm += data[i][0]
	}

	maxSumm -= minSumm
	var left int = n - n/2
	var sred int = maxSumm / left

	if data[n-1][0] < sred {
		fmt.Print(data[n-1][0])
	} else {
		fmt.Print(sred)
	}
}
