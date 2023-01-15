package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type adList struct {
	l     map[int][]int
	paths [][]int
	path  []int
}

func initAdlist(v_count int) adList {
	return adList{
		l:     make(map[int][]int, v_count+1),
		paths: make([][]int, 0, 2),
		path:  make([]int, 0, 2),
	}
}
func (al *adList) appendPaths() {
	tmp := make([]int, len(al.path))
	copy(tmp, al.path)
	al.paths = append(al.paths, tmp)
}
func (al adList) printWithSpace() {
	for _, l := range al.l {
		for j, v := range l {
			if j >= 1 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", v)
		}
		fmt.Printf("\n")
	}
}
func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	s := scanInt()
	t := scanInt()
	// fmt.Printf("%d %d %d\n", n, s, t)

	for i := 0; i < n; i++ {
		v := scanInt()
		for i := 0; i < v; i++ {
			a := scanInt()

		}
	}
}

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}
func scanText() string {
	sc.Scan()
	return sc.Text()
}
