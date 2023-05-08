package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fLine := getIntSlice(in)
	var n, x, t, res int = fLine[0], fLine[1], fLine[2], 0
	fLine = getIntSlice(in)

	var sculps [][2]int = make([][2]int, n)
	for i := 0; i < n; i++ {
		sculps[i][0] = abs(fLine[i], x)
		sculps[i][1] = i + 1
	}

	// fmt.Print(n, x, t, sculps)

	sort.Slice(sculps, func(a, b int) bool {
		return sculps[a][0] < sculps[b][0]
	})

	for i := 0; i < n; i++ {
		if sculps[i][0] <= t {
			res++
			t -= sculps[i][0]
			sculps[i][0] = 0
			// resLine += strconv.Itoa(sculps[i][1]) + " "
		}
	}

	var resLine []string = make([]string, res)
	var z = 0

	for i := 0; i < n; i++ {
		if sculps[i][0] == 0 {
			resLine[z] = strconv.Itoa(sculps[i][1])
			z++
		} else {
			break
		}
	}

	out.WriteString(strconv.Itoa(res) + "\n" + strings.Join(resLine, " "))
	out.Flush()
}

func getIntSlice(sc *bufio.Reader) []int {
	stri, _ := sc.ReadString('\n')
	slice := strings.Split(stri[:len(stri)-1], " ")
	var res []int = make([]int, len(slice))

	for i := 0; i < len(slice); i++ {
		res[i], _ = strconv.Atoi(slice[i])
	}

	return res
}

func abs(a, b int) int {
	if a-b >= 0 {
		return a - b
	}
	return b - a
}
