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
	moveCnt := strToInt(line1[2])

	//fmt.Printf("X Y moveCnt %d %d %d\n", X, Y, moveCnt)

	distance := 0
	// moveCntと同じ数までループさせるため初期値１にする
	totalDistance := 0
	i := 0
	for totalDistance < moveCnt {
		//fmt.Printf("移動前 %d %d\n", X, Y)

		isMoveRight := i%4 == 0
		isMoveUp := i%4 == 1
		isMoveLeft := i%4 == 2
		isMoveUnder := i%4 == 3

		if i%2 == 0 {
			distance += 1
		}

		// Nを超える移動数の場合丸める
		if totalDistance+distance < moveCnt {
			totalDistance += distance
		} else {
			distance = moveCnt - totalDistance
			totalDistance += distance
		}

		//fmt.Printf("i totalDistance distance moveCnt %d %d %d %d\n", i, totalDistance, distance, moveCnt)

		if isMoveRight {
			X += distance
		} else if isMoveUnder {
			Y -= distance
		} else if isMoveLeft {
			X -= distance
		} else if isMoveUp {
			Y += distance
		}

		// 以下で2回同じ距離を移動する規則性を実現する

		i++
		//fmt.Printf("移動後! %d %d %d %d %d %d\n", X, Y, totalDistance, distance, moveCnt, i)
	}
	fmt.Printf("%d %d\n", X, Y)

}

func strToInt(s string) int {
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}
