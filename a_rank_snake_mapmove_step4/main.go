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
	line1 := strings.Split(readLine(), " ")

	H := strToInt(line1[0])
	//W := strToInt(line1[1])
	sy := strToInt(line1[2])
	sx := strToInt(line1[3])
	N := strToInt(line1[4])
	//fmt.Printf("H W sy sx N %d %d %d %d %d\n", H, W, sy, sx, N)

	area := make([][]string, 0)
	for i := 0; i < H; i++ {
		si := strings.Split(readLine(), "")
		area = append(area, si)
	}
	//fmt.Printf("%v\n", area)

	person := Person{
		y:    sy,
		x:    sx,
		di:   "N",
		area: area,
	}

	for i := 0; i < N; i++ {
		di := strings.Split(readLine(), " ")
		//fmt.Println(di)
		canMove := person.move(di[0], strToInt(di[1]))
		if !canMove {
			break
		}
	}
}

func (p *Person) canMove() bool {
	H := len(p.area)
	W := len(p.area[0])
	// 現在地から前に移動できない時
	if (p.di == "N" && p.y == 0) ||
		(p.di == "E" && p.x == W-1) ||
		(p.di == "S" && p.y == H-1) ||
		(p.di == "W" && p.x == 0) {
		return false
	}

	if (p.di == "N" && p.area[p.y-1][p.x] == "#") ||
		(p.di == "E" && p.area[p.y][p.x+1] == "#") ||
		(p.di == "S" && p.area[p.y+1][p.x] == "#") ||
		(p.di == "W" && p.area[p.y][p.x-1] == "#") {
		return false
	}
	return true
}

func (p *Person) moveForward() {
	if p.di == "N" {
		p.y -= 1
	} else if p.di == "E" {
		p.x += 1
	} else if p.di == "S" {
		p.y += 1
	} else if p.di == "W" {
		p.x -= 1
	}
}

func (p *Person) changeDi(moveDi string) {
	if moveDi == "L" {
		if p.di == "N" {
			p.di = "W"
		} else if p.di == "E" {
			p.di = "N"
		} else if p.di == "S" {
			p.di = "E"
		} else if p.di == "W" {
			p.di = "S"
		}
	}

	if moveDi == "R" {
		if p.di == "N" {
			p.di = "E"
		} else if p.di == "E" {
			p.di = "S"
		} else if p.di == "S" {
			p.di = "W"
		} else if p.di == "W" {
			p.di = "N"
		}
	}

}

// 受け取った方向に向き変えて
// 繰り返し
// 	移動可能か確認してOK
//		前に進む
//		現在の座標 を出力
// 	NGだったら
//		Stop を出力
func (p *Person) move(moveDi string, forwardNum int) bool {
	//fmt.Printf("%s に %d 回進みます\n", moveDi, forwardNum)
	p.changeDi(moveDi)
	canMove := true
	for i := 0; i < forwardNum; i++ {
		if p.canMove() {
			p.moveForward()
		} else {
			canMove = false
		}
	}
	if canMove {
		fmt.Printf("%d %d\n", p.y, p.x)
		return true
	} else {
		fmt.Printf("%d %d\n", p.y, p.x)
		fmt.Println("Stop")
		return false
	}
}

type Person struct {
	y    int
	x    int
	di   string // N or E or S or W
	area [][]string
}

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	r, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	return string(r)
}

func strToInt(s string) int {
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}
