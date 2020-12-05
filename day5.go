package main

import (
	"fmt"
	"io/ioutil"
	"sort"

	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getID(y string) int {
	r := ""
	c := ""
	for i := range y {
		switch y[i] {
		case 'B':
			r += "1"
		case 'F':
			r += "0"
		case 'L':
			c += "0"
		case 'R':
			c += "1"
		}
	}

	ri, _ := strconv.ParseInt(r, 2, 8)
	ci, _ := strconv.ParseInt(c, 2, 4)
	return int(ri*8 + ci)
}

func main() {
	tinput := time.Now()

	b, err := ioutil.ReadFile("d5.data")
	check(err)
	sref := strings.Split(string(b), "\n")

	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)
	tinput = time.Now()
	v := []int{}
	for _, y := range sref {
		v = append(v, getID(y))
	}
	sort.Ints(v)
	max := v[len(v)-1]
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println(max)
	tinput = time.Now()
	a := 0
	for i := range v {
		if v[i+1]-v[i] == 2 {
			a = v[i+1] - 1
			break
		}
	}
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println(a)
}
