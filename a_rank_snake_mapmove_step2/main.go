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
	d := line[4]
	m := line[5]
	//fmt.Printf("H W sy sx m: %d %d %d %d %s\n", H, W, sy, sx, m)
	person := Person{X: sx, Y: sy, direction: d}

	//fmt.Printf("person %v\n", person)

	var maps [][]string
	for i := 0; i < H; i++ {
		Si := strings.Split(nextLine(), "")
		maps = append(maps, Si)
	}
	//fmt.Printf("%v\n", maps)
	//fmt.Printf("%s\n", maps[sy][sx])
	person.move(m)

	//fmt.Printf("moved person %v\n", person)

	if person.Y < 0 ||
		person.Y > H-1 ||
		person.X < 0 ||
		person.X > W-1 ||
		maps[person.Y][person.X] == "#" {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

type Person struct {
	X         int
	Y         int
	direction string
}

func (p *Person) moveLeft() {
	if p.direction == "N" {
		p.X -= 1
	} else if p.direction == "E" {
		p.Y -= 1
	} else if p.direction == "S" {
		p.X += 1
	} else if p.direction == "W" {
		p.Y += 1
	}
}

func (p *Person) moveRight() {
	if p.direction == "N" {
		p.X += 1
	} else if p.direction == "E" {
		p.Y += 1
	} else if p.direction == "S" {
		p.X -= 1
	} else if p.direction == "W" {
		p.Y -= 1
	}
}

func (p *Person) move(m string) {
	if m == "L" {
		p.moveLeft()
	} else if m == "R" {
		p.moveRight()
	}
}

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
