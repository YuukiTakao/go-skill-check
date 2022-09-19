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
	//fmt.Printf("H: %d W: %d N: %d\n", H, W, N)

	var yi []string
	for i := 0; i < H; i++ {
		scanner.Scan()
		row := strings.Split(scanner.Text(), "")
		for _, s := range row {
			yi = append(yi, s)
		}
	}
	//fmt.Printf("yi: %s\n", yi)

	var ni []int
	for i := 0; i < N; i++ {
		scanner.Scan()
		row := strings.Split(scanner.Text(), " ")

		ny := 0
		tmpy := strToInt(row[0])
		if tmpy == 0 {
			ny = tmpy
		} else {
			ny = tmpy * W
		}
		nx := strToInt(row[1])

		//fmt.Printf("ny: %d nx %d\n", ny, nx)
		ni = append(ni, ny+nx)
	}
	//fmt.Printf("ni: %s\n", ni)

	for i := 0; i < N; i++ {
		yi[ni[i]] = "#"
	}
	for i := 0; i < H; i++ {
		row := ""
		for j := 0; j < W; j++ {

			idx := 0
			if i == 0 {
				idx = i + j
			} else {
				idx = i*W + j
			}
			//fmt.Printf("idx: %d\n", idx)
			row += yi[idx]
		}
		fmt.Println(row)
	}
}

func strToInt(s string) int {
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}
