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
	M := strToInt(line1[1])
	//fmt.Printf("N: %d M: %d \n", N, M)

	scanner.Scan()
	data := strings.Split(scanner.Text(), " ")
	fmt.Println(data[M-1])
}

func strToInt(s string) int {
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}
