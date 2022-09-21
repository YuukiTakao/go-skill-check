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

	Y := strToInt(line1[0])
	X := strToInt(line1[1])
	N := strToInt(line1[2])

	//fmt.Printf("Y X N %d %d %d\n", Y, X, N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		move := scanner.Text()
		//fmt.Println(move)
		if move == "N" {
			Y -= 1
			fmt.Printf("%d %d\n", Y, X)
		} else if move == "S" {
			Y += 1
			fmt.Printf("%d %d\n", Y, X)
		} else if move == "E" {
			X += 1
			fmt.Printf("%d %d\n", Y, X)
		} else if move == "W" {
			X -= 1
			fmt.Printf("%d %d\n", Y, X)
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
