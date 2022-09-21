package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N := strToInt(scanner.Text())
	//fmt.Printf("N %d\n", N)

	scanner.Scan()
	A := strings.Split(scanner.Text(), " ")

	am := make(map[int]int, N)
	for i, a := range A {
		am[i+1] = strToInt(a)
	}

	scanner.Scan()
	Q := strToInt(scanner.Text())

	scanner.Scan()
	B := strings.Split(scanner.Text(), " ")
	bm := make(map[int]int, Q)
	for i, b := range B {
		bm[i+1] = strToInt(b)
	}
	//fmt.Printf("am %d\n", am)
	//fmt.Printf("bm %d\n", bm)

	for i := 1; i <= Q; i++ {
		fmt.Println(am[bm[i]])
	}

}

func strToInt(s string) int {
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}
