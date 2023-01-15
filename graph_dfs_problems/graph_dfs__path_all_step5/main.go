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
func inSlice(slice []int, target int) bool {
	for _, num := range slice {
		if num == target {
			return true
		}
	}
	return false
}
func (al *adList) walkGraphDfsStartToGoalWithTransit(start, goal, transit int, isUnique bool) [][]int {
	al.path = append(al.path, start)
	al.dfs(start, goal, transit, isUnique)
	return al.paths
}
func (al *adList) dfs(position, goal, transit int, isUnique bool) {
	for _, v := range al.l[position-1] {
		if !inSlice(al.path, v) || !isUnique {
			al.path = append(al.path, v)
			if v == goal && inSlice(al.path, transit) {
				al.appendPaths()
			} else {
				al.dfs(v, goal, transit, isUnique)
			}
			al.path = al.path[:len(al.path)-1]
		}
	}
}
func printWithSpace(s2d [][]int) {
	for _, s := range s2d {
		for j, v := range s {
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
	p := scanInt()
	// fmt.Printf("%d %d %d %d\n", n, s, t, p)

	al := initAdlist(n)
	for i := 0; i < n; i++ {
		v := scanInt()
		a := make([]int, v)
		for j := 0; j < v; j++ {
			a[j] = scanInt()
		}
		al.l[i] = a
	}
	paths := al.walkGraphDfsStartToGoalWithTransit(s, t, p, true)
	fmt.Printf("%d\n", len(paths))
	printWithSpace(paths)
	// adlist.printWithSpace()

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
