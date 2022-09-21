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
	X := strToInt(line1[0])
	Y := strToInt(line1[1])
	N := strToInt(line1[2])
	//fmt.Printf("X Y N %d %d %d\n", X, Y, N)

	person := Person{X: X, Y: Y, direction: "N"}

	for i := 0; i < N; i++ {

		scanner.Scan()
		di := scanner.Text()
		//fmt.Printf("di %s i %d\n", di, i)
		person.move(di)
		fmt.Printf("%d %d\n", person.X, person.Y)
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
		p.direction = "W"
	} else if p.direction == "W" {
		p.Y += 1
		p.direction = "S"
	} else if p.direction == "S" {
		p.X += 1
		p.direction = "E"
	} else if p.direction == "E" {
		p.Y -= 1
		p.direction = "N"
	}
}

func (p *Person) moveRight() {
	if p.direction == "N" {
		p.X += 1
		p.direction = "E"
	} else if p.direction == "W" {
		p.Y -= 1
		p.direction = "N"
	} else if p.direction == "S" {
		p.X -= 1
		p.direction = "W"
	} else if p.direction == "E" {
		p.Y += 1
		p.direction = "S"
	}
}

func (p *Person) move(moveDi string) (int, int) {
	if moveDi == "L" {
		p.moveLeft()
	} else if moveDi == "R" {
		p.moveRight()
	}
	return p.X, p.Y
}

func strToInt(s string) int {
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}
