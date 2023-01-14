package main

import (
	"fmt"
)

func graph_walk_dfs(n, s, k int, adlist *map[int][]int) [][]int {
	routes := make([][]int, 0, 2)
	walk := make([]int, 0, 2)
	walk = append(walk, s)
	dfs(k, s, &routes, &walk, adlist)
	return routes
}

func dfs(k, position int, routes *[][]int, walk *[]int, adlist *map[int][]int) {
	if k == 0 {
		tmp := make([]int, len(*walk))
		copy(tmp, *walk)
		*routes = append(*routes, tmp)
	} else {
		tadlist := *adlist
		twalk := *walk
		for _, v := range tadlist[position-1] {
			twalk = append(twalk, v)
			dfs(k-1, v, routes, &twalk, adlist)
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

func main() {
	var n, s, k int
	fmt.Scanf("%d %d %d", &n, &s, &k)

	adlist := make(map[int][]int, n+1)

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
	routes := graph_walk_dfs(n, s, k, &adlist)

	fmt.Printf("%d\n", len(routes))
	print_2dslice_with_space(routes)
}
