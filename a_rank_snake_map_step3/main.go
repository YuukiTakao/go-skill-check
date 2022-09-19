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
	//fmt.Printf("H: %d W: %d\n", H, W)

	for i := 0; i < H; i++ {
		scanner.Scan()
		sl := strings.Split(scanner.Text(), "")
		for j, _ := range sl {
			//fmt.Println(s)
			if j == 0 {
				if sl[j+1] == "#" {
					fmt.Printf("%d %d\n", i, j)
				}
			} else if len(sl)-1 == j {
				if sl[j-1] == "#" {
					fmt.Printf("%d %d\n", i, j)
				}
			} else {
				if sl[j-1] == "#" && sl[j+1] == "#" {
					fmt.Printf("%d %d\n", i, j)
				}
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
