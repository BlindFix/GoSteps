package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n, m, k int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &n, &m, &k)

	var cluster int = n / m
	var remainingJuns int = n % m
	var timeAdditional int = int(math.Ceil(float64(remainingJuns) * float64(k) / float64(m)))
	fmt.Print(cluster*k + timeAdditional)
}
