package main

import (
	"io/ioutil"
	"log"

	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func mAtoi(v string) int {
	c, _ := strconv.Atoi(v)
	return c
}

func slope(r int, d int, forest []string) int {
	mapl := len(forest[0])
	tco := 0
	x := 0
	y := 0
	for y < len(forest) {
		if forest[y][x] == 35 {
			tco++
		}
		x = (x + r) % mapl
		y += d
	}
	return tco
}

func main() {
	tinput := time.Now()

	b, err := ioutil.ReadFile("d3.data")
	check(err)
	forest := strings.Split(string(b), "\n")

	inputtelapsed := time.Since(tinput)
	log.Println("Read data:", inputtelapsed)
	tinput = time.Now()

	// 46 == .
	// 35 == #
	mapl := len(forest[0])
	x := 0
	tc := 0
	for y := range forest {
		if forest[y][x] == 35 {
			tc++
		}
		x = (x + 3) % mapl
	}
	inputtelapsed = time.Since(tinput)
	log.Println("Time p1:", inputtelapsed)
	log.Println(tc)

	tinput = time.Now()
	ml := slope(1, 1, forest) * slope(3, 1, forest) * slope(5, 1, forest) * slope(7, 1, forest) * slope(1, 2, forest)

	inputtelapsed = time.Since(tinput)
	log.Println("Time p2:", inputtelapsed)
	log.Println(ml)
}
