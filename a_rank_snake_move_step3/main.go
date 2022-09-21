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
	D := line1[2]

	//fmt.Printf("Y X D %d %d %s\n", Y, X, D)

	scanner.Scan()
	move := scanner.Text()

	if D == "N" {
		if move == "R" {
			X += 1
		} else if move == "L" {
			X -= 1
		}
	} else if D == "E" {
		if move == "R" {
			Y += 1
		} else if move == "L" {
			Y -= 1
		}
	} else if D == "S" {
		if move == "R" {
			X -= 1
		} else if move == "L" {
			X += 1
		}
	} else if D == "W" {
		if move == "R" {
			Y -= 1
		} else if move == "L" {
			Y += 1
		}
	}
	fmt.Printf("%d %d", Y, X)
}

func strToInt(s string) int {
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}
