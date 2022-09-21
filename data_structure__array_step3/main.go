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
	line := strings.Split(scanner.Text(), " ")

	min := 0
	max := 0
	for _, s := range line {
		n := strToInt(s)
		if min == 0 {
			min = n
		}
		if max == 0 {
			max = n
		}
		if min > n {
			min = n
		}
		if max < n {
			max = n
		}
	}
	fmt.Printf("%d", max-min)
}

func strToInt(s string) int {
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}
