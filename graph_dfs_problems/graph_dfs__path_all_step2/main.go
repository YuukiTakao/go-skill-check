package main

import (
	"fmt"
)

type adList struct {
	l     map[int][]int
	paths [][]int
	path  []int
}

func init_adlist(v_count int) adList {
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

func (al adList) print_with_space() {
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
func (al *adList) walkGraphDfsInKMoves(start, k int, isUnique bool) [][]int {
	al.path = append(al.path, start)
	al.dfs(start, k, isUnique)
	return al.paths
}
func (al *adList) dfs(position, k int, isUnique bool) {
	if k == 0 {
		al.appendPaths()
	} else {
		for _, v := range al.l[position-1] {
			if !inSlice(al.path, v) || !isUnique {
				al.path = append(al.path, v)
				al.dfs(v, k-1, isUnique)
				al.path = al.path[:len(al.path)-1]
			}
		}
	}
}
func print_2dslice_with_space(tds [][]int) {
	for _, s := range tds {
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
	var n, s, k int
	fmt.Scan(&n, &s, &k)

	adlist := init_adlist(n)

	var v, a int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &v)
		as := make([]int, 0, v)

		for j := 0; j < v; j++ {
			fmt.Scanf("%d", &a)
			as = append(as, a)
		}
		adlist.l[i] = as
	}
	paths := adlist.walkGraphDfsInKMoves(s, k, false)

	fmt.Printf("%d\n", len(paths))
	print_2dslice_with_space(paths)
}
