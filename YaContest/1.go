package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var n = getIntSlice(in)[0]
	var keySym = getIntSlice(in)
	var keyRow = getIntSlice(in)
	var k = getIntSlice(in)[0]
	var text = getIntSlice(in)

	var keyMap = make(map[int]int)
	for i := 0; i < n; i++ {
		keyMap[keySym[i]] = keyRow[i]
	}

	var res = 0
	for i := 1; i < k; i++ {
		if keyMap[text[i-1]] != keyMap[text[i]] {
			res++
		}
	}

	out.WriteString(strconv.Itoa(res))
	out.Flush()
}

func getIntSlice(sc *bufio.Reader) []int {
	stri, _ := sc.ReadString('\n')
	stri = strings.Trim(stri, "\n")
	slice := strings.Split(stri, " ")
	var res []int = make([]int, len(slice))

	for i := 0; i < len(slice); i++ {
		res[i], _ = strconv.Atoi(slice[i])
	}

	return res
}
