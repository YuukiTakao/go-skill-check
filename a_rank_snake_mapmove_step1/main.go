package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	line := strings.Split(nextLine(), " ")

	H := strToInt(line[0])
	W := strToInt(line[1])
	sy := strToInt(line[2])
	sx := strToInt(line[3])
	m := line[4]
	//fmt.Printf("H W sy sx m: %d %d %d %d %s\n", H, W, sy, sx, m)

	var maps [][]string
	for i := 0; i < H; i++ {
		Si := strings.Split(nextLine(), "")
		maps = append(maps, Si)
	}
	//fmt.Printf("%v\n", maps)
	//fmt.Printf("%s\n", maps[sy][sx])

	if m == "N" {
		sy -= 1
	} else if m == "E" {
		sx += 1
	} else if m == "S" {
		sy += 1
	} else if m == "W" {
		sx -= 1
	}
	//fmt.Printf("%d %d\n", sy, sx)

	if sy < 0 || sy > H-1 || sx < 0 || sx > W-1 {
		fmt.Println("No")
	} else if maps[sy][sx] == "#" {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

//type Person struct {
//	X int
//	Y int
//}

func nextLine() string {
	scanner.Scan()
	return scanner.Text()
}

func strToInt(s string) int {
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}
