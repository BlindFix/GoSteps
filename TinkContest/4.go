package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &n)
	var arr []int = make([]int, n)

	var countMap map[int]int = make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
		countMap[arr[i]] = 0
	}

	var avg, avgCount, max, maxCount, oneCount, difElemCount int = 1, 1, 2, 0, 1, 1
	countMap[arr[0]]++

	var res int = 1
	for i := 1; i < n; i++ {
		countMap[arr[i]]++
		if countMap[arr[i]] == 1 {
			oneCount++
			difElemCount++
		} else if countMap[arr[i]] == 2 {
			oneCount--
		}

		if countMap[arr[i]] == avg {
			avgCount++
		} else if countMap[arr[i]] == max {
			maxCount++
			avgCount--
			if maxCount > 1 {
				avg = max
				avgCount = maxCount
				max++
				maxCount = 0
			}
		} else if countMap[arr[i]] > max {
			max = countMap[arr[i]]
			avgCount = maxCount
			maxCount = 1
			avg = max - 1
		}
		if (avgCount == difElemCount && avg == 1) || (avgCount == difElemCount-1 && (maxCount == 1 || oneCount > 0)) {
			res = i + 1
		}

		// fmt.Printf("iter %d, elem %d, count %d, %v, avg %d, avgcnt %d, max %d, maxCnt %d, res %d\n", i, arr[i], countMap[arr[i]], countMap, avg, avgCount, max, maxCount, res)
	}

	fmt.Print(res)
}
