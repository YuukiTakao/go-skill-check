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
	W := strToInt(line1[1])
	N := strToInt(line1[2])
	//fmt.Printf("H: %d W: %d N: %d", H, W, N)

	var yi []string
	for i := 0; i < H; i++ {
		scanner.Scan()

		var line []string
		line = strings.Split(scanner.Text(), "")
		for _, s := range line {
			yi = append(yi, s)
		}

	}
	//fmt.Printf("yi: %s\n", yi)

	var ni []int
	for i := 0; i < N; i++ {
		scanner.Scan()
		tmp := strings.Split(scanner.Text(), " ")

		tmpy := strToInt(tmp[0])
		ny := 0
		if tmpy == 0 {
			ny = tmpy
		} else {
			ny = tmpy * W
		}
		nx := strToInt(tmp[1])
		//fmt.Printf("ny: %d  nx: %d\n", ny, nx)
		ni = append(ni, ny+nx)
	}
	//fmt.Printf("ni: %s \n", ni)

	for i := 0; i < N; i++ {
		fmt.Println(yi[ni[i]])
	}
}

func strToInt(s string) int {
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}
