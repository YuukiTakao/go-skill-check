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
	line1 := strings.Split(scanner.Text(), " ")

	//N := strToInt(line1[0])
	Q := strToInt(line1[1])
	//fmt.Printf("N Q %d %d\n", N, Q)

	scanner.Scan()
	an := strings.Split(scanner.Text(), " ")

	//fmt.Printf("an %s\n", an)

	//var result []int
	for i := 0; i < Q; i++ {
		scanner.Scan()
		qi := strings.Split(scanner.Text(), " ")

		//fmt.Printf("qi qi[0] %s %s\n", qi, qi[0])

		if qi[0] == "0" {
			an = append(an, qi[1])
		} else if qi[0] == "1" {
			an = an[:len(an)-1]
		} else if qi[0] == "2" {
			fmt.Printf("%s\n", strings.Join(an, " "))
		}
	}
}

func strToInt(s string) int {
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}
