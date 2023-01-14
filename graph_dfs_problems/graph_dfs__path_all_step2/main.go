package main

import (
	"fmt"
)

func graph_walk_dfs_k_moves(start, k int, adlist *map[int][]int) [][]int {
	routes := make([][]int, 0, 2)
	walk := make([]int, 0, 2)
	walk = append(walk, start)
	dfs(start, k, &routes, &walk, adlist)
	return routes
}
func dfs(position, k int, routes *[][]int, walk *[]int, adlist *map[int][]int) {
	if k == 0 {
		tmp := make([]int, len(*walk))
		copy(tmp, *walk)
		*routes = append(*routes, tmp)
	} else {
		tadlist := *adlist
		twalk := *walk
		for _, v := range tadlist[position-1] {
			twalk = append(twalk, v)
			dfs(v, k-1, routes, &twalk, adlist)
			twalk = twalk[:len(twalk)-1]
		}
	}
}

func print_2dslice_with_space(s2d [][]int) {
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

func init_adlist(v_count int) map[int][]int {
	return make(map[int][]int, v_count+1)
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
		adlist[i] = as
	}
	routes := graph_walk_dfs_k_moves(s, k, &adlist)

	fmt.Printf("%d\n", len(routes))
	print_2dslice_with_space(routes)
}
