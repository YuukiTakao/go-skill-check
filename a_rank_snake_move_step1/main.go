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

	H := strToInt(line1[0])
	//W := strToInt(line1[1])

	//var flatS []string
	for i := 0; i < H; i++ {
		scanner.Scan()
		sl := strings.Split(scanner.Text(), "")

		for j, s := range sl {
			//flatS = append(flatS, s)
			if s == "#" {
				fmt.Printf("%d %d", i, j)
			}
		}
	}
	//fmt.Printf("flatS %s", flatS)
}

func strToInt(s string) int {
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}
