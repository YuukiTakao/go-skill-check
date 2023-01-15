package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type adList struct {
	l      map[int][]int
	routes [][]int
	walk   []int
}

func (al *adList) graph_walk_dfs_k_moves(start, k int) [][]int {
	al.walk = append(al.walk, start)
	al.dfs(start, k)
	return al.routes
}
func (al *adList) dfs(position, k int) {
	if k == 0 {
		tmp := make([]int, len(al.walk))
		copy(tmp, al.walk)
		al.routes = append(al.routes, tmp)
	} else {
		// twalk := *walk
		for _, v := range al.l[position-1] {
			al.walk = append(al.walk, v)
			al.dfs(v, k-1)
			al.walk = al.walk[:len(al.walk)-1]
		}
	}
}
func init_adlist(v_count int) adList {
	return adList{
		l:      make(map[int][]int, v_count+1),
		routes: make([][]int, 0, 2),
		walk:   make([]int, 0, 2),
	}
}

type tdiSlice struct {
	sl [][]int
}

func init_tdislice(count int) tdiSlice {
	return tdiSlice{
		sl: make([][]int, count),
	}
}
func (t tdiSlice) print_2dslice_with_space(s2d [][]int) {
	for _, s := range t.sl {
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
	k := scanInt()
	fmt.Printf("%d %d %d\n", n, s, k)

	adlist := init_adlist(n)
	for i := 0; i < n; i++ {
		v := scanInt()
		a := make([]int, v)
		for i := 0; i < v; i++ {
			a[i] = scanInt()
		}
		adlist.l[i] = a
	}

	adlist.print_with_space()
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
