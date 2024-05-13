package main

import "fmt"

func main() {
	mapa := map[int]int{}

	for i := 0; i < 3; i++ {
		mapa[1] += 1
	}

	fmt.Println(mapa)
}
