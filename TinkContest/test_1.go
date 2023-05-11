package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &a, &b)
	fmt.Print(a + b)
}
