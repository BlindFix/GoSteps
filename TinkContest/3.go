package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var s string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &n, &s)

	mapSlice := make(map[rune]int)
	mapSlice[rune('a')] = 0
	mapSlice[rune('b')] = 0
	mapSlice[rune('c')] = 0
	mapSlice[rune('d')] = 0

	var LP, RP, minLen, howMany int = 0, 0, n + 1, 0
	var ok bool = true
	const needed int = 4

	for RP < n {
		mapSlice[rune(s[RP])]++
		if mapSlice[rune(s[RP])] == 1 {
			howMany++

			if howMany == needed {
				break
			}
		}

		RP++
	}

	if howMany != needed {
		fmt.Print(-1)
	} else {
		for ok && LP < n {
			minLen = RP - LP + 1
			mapSlice[rune(s[LP])]--
			if mapSlice[rune(s[LP])] == 0 {
				howMany--
			}
			LP++

			for howMany < needed {
				RP++
				if RP < n {
					mapSlice[rune(s[RP])]++
					if mapSlice[rune(s[RP])] == 1 {
						howMany++
					}
				} else {
					ok = false
					break
				}
			}
			if ok && RP-LP+1 < minLen {
				minLen = RP - LP + 1
			}
		}

		fmt.Print(minLen)
	}
}
