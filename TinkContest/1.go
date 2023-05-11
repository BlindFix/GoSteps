package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x1, x2, x3, x4 int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &x1, &x2, &x3, &x4)

	if (x1 <= x2 && x2 <= x3 && x3 <= x4) || (x1 >= x2 && x2 >= x3 && x3 >= x4) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
