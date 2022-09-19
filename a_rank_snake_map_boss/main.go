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
	W := strToInt(line1[0])
	//fmt.Printf("H %d W %d\n", H, W)

	var flatS []string
	for i := 0; i < H; i++ {
		scanner.Scan()
		sl := strings.Split(scanner.Text(), "")
		for _, s := range sl {
			flatS = append(flatS, s)
		}
	}
	//fmt.Printf("flatS %s \n", flatS)

	for i, _ := range flatS {
		//fmt.Printf("i %d s %s\n", i, s)
		y := i / H
		x := i % H
		if sorroundSharp(flatS, H, W, i) {
			fmt.Printf("%d %d\n", y, x)
		}
	}
}

func leftExist(x int) bool {
	return x-1 >= 0
}

func rightExist(x int, W int) bool {
	return x+1 < W
}

func upExist(y int) bool {
	return y-1 >= 0
}

func underExist(y int, H int) bool {
	return y+1 < H
}

func isSharp(cells []string, idx int) bool {
	return cells[idx] == "#"
}

// 該当セルが存在する場合にのみ#かどうかチェックする
func sorroundSharp(cells []string, H int, W int, idx int) bool {

	y := idx / H
	x := idx % W

	if leftExist(x) {
		if !isSharp(cells, idx-1) {
			return false
		}
	}
	if rightExist(x, W) {
		if !isSharp(cells, idx+1) {
			return false
		}
	}
	if upExist(y) {
		if !isSharp(cells, idx-W) {
			return false
		}
	}
	if underExist(y, H) {
		if !isSharp(cells, idx+W) {
			return false
		}
	}

	return true
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
