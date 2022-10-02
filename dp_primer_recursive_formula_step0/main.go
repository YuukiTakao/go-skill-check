package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x, d, k int
	fmt.Fscan(reader, &x, &d, &k)
	//fmt.Printf("%d %d %d\n", x, d, k)

	a := make([]int, k)
	a[0] = x
	for i := 1; i < k; i++ {
		a[i] = a[i-1] + d
	}

	fmt.Printf("%d\n", a[k-1])
}

var reader = bufio.NewReader(os.Stdin)
