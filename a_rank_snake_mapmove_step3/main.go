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
	W := strToInt(line1[1])
	sy := strToInt(line1[2])
	sx := strToInt(line1[3])
	N := strToInt(line1[4])

	//fmt.Printf("H %d W %d sy %d sx %d N %d\n", H, W, sy, sx, N)

	area := make([][]string, 0)
	for i := 0; i < H; i++ {
		si := strings.Split(readLine(), "")
		//fmt.Printf("si %v\n", si)
		area = append(area, si)
	}
	//fmt.Printf("%v\n", area)

	person := Person{
		y:    sy,
		x:    sx,
		di:   "N",
		area: area,
	}
	//fmt.Printf("%v\n", person)

	for i := 0; i < N; i++ {
		di := readLine()
		//fmt.Printf("%s\n", di)
		canMove := person.move(di, H, W)

		if !canMove {
			break
		}

	}

}

func (p *Person) rightUpdate() {
	//before := p.di
	if p.di == "N" {
		p.di = "E"
	} else if p.di == "E" {
		p.di = "S"
	} else if p.di == "S" {
		p.di = "W"
	} else if p.di == "W" {
		p.di = "N"
	}
	//fmt.Printf("%s から %s に向き変えました \n", before, p.di)
}

func (p *Person) leftUpdate() {
	//before := p.di
	if p.di == "N" {
		p.di = "W"
	} else if p.di == "E" {
		p.di = "N"
	} else if p.di == "S" {
		p.di = "E"
	} else if p.di == "W" {
		p.di = "S"
	}
	//fmt.Printf("%s から %s に向き変えました \n", before, p.di)
}

func (p *Person) directionUpdate(moveDi string) {
	//fmt.Printf("方向を %s に変えます \n", moveDi)
	if moveDi == "L" {
		p.leftUpdate()
	} else if moveDi == "R" {
		p.rightUpdate()
	}
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

func (p *Person) canMove(H int, W int) bool {
	// マップ範囲外
	if (p.di == "N" && p.y == 0) ||
		(p.di == "E" && p.x == W-1) ||
		(p.di == "S" && p.y == H-1) ||
		(p.di == "W" && p.x == 0) {
		return false
	}
	// 次のマスが壁になっている
	if (p.di == "N" && p.area[p.y-1][p.x] == "#") ||
		(p.di == "E" && p.area[p.y][p.x+1] == "#") ||
		(p.di == "S" && p.area[p.y+1][p.x] == "#") ||
		(p.di == "W" && p.area[p.y][p.x-1] == "#") {
		return false
	}
	return true
}

func (p *Person) move(moveDi string, H int, W int) bool {
	p.directionUpdate(moveDi)
	if p.canMove(H, W) {
		p.moveForward()
		fmt.Printf("%d %d\n", p.y, p.x)

		return true
	} else {
		fmt.Println("Stop")

		return false
	}

}

type Person struct {
	y    int
	x    int
	di   string
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
