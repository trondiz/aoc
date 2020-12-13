package main

import (
	"io/ioutil"
	"log"

	"strconv"
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

type pos struct {
	x, y int
}

func main() {
	tinput := time.Now()

	dir, err := ioutil.ReadFile("d3.data")
	check(err)

	inputtelapsed := time.Since(tinput)
	log.Println("Read data:", inputtelapsed)
	tinput = time.Now()
	g := map[pos]int{pos{x: 0, y: 0}: 1}

	c := 0
	x := 0
	y := 0
	for _, d := range dir {
		switch d {
		case 'v':
			y--
			g[pos{x: x, y: y}]++
		case '>':
			x++
			g[pos{x: x, y: y}]++
		case '^':
			y++
			g[pos{x: x, y: y}]++
		case '<':
			x--
			g[pos{x: x, y: y}]++
		}

	}
	c = len(g)
	inputtelapsed = time.Since(tinput)
	log.Println("Time p1:", inputtelapsed)
	log.Println(c)
	tinput = time.Now()

	g = map[pos]int{pos{x: 0, y: 0}: 1}

	c = 0
	x = 0
	y = 0
	rx := 0
	ry := 0
	for i, d := range dir {
		if i%2 == 0 {
			switch d {
			case 'v':
				y--
				g[pos{x: x, y: y}]++
			case '>':
				x++
				g[pos{x: x, y: y}]++
			case '^':
				y++
				g[pos{x: x, y: y}]++
			case '<':
				x--
				g[pos{x: x, y: y}]++
			}
		} else {
			switch d {
			case 'v':
				ry--
				g[pos{x: rx, y: ry}]++
			case '>':
				rx++
				g[pos{x: rx, y: ry}]++
			case '^':
				ry++
				g[pos{x: rx, y: ry}]++
			case '<':
				rx--
				g[pos{x: rx, y: ry}]++
			}
		}

	}
	c = len(g)

	inputtelapsed = time.Since(tinput)
	log.Println("Time p2:", inputtelapsed)
	log.Println(c)
}
