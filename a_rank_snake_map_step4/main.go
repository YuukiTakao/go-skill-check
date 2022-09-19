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
	//fmt.Printf("H: %d\n", H)

	var flatS []string
	for i := 0; i < H; i++ {
		scanner.Scan()
		sl := strings.Split(scanner.Text(), "")
		for _, s := range sl {
			//fmt.Printf("s: %s i*(j+1)+W: %d \n", s, (j+1)*i+W)
			flatS = append(flatS, s)
		}
	}
	//fmt.Printf("flatS: %s\n", flatS)

	for i, _ := range flatS {
		//fmt.Printf("i: %d i/W+1: %d iaW: %d\n", i, i/W+1, i%W)
		if i < W {
			if flatS[i+W] == "#" {
				fmt.Printf("%d %d\n", i/W, i%W)
			}
		} else if i/W+1 == H {
			if flatS[i-W] == "#" {
				fmt.Printf("%d %d\n", i/W, i%W)
			}
		} else {
			if flatS[i-W] == "#" && flatS[i+W] == "#" {
				fmt.Printf("%d %d\n", i/W, i%W)
			}
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
