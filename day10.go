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

func mAtoi(v string) int {
	c, _ := strconv.Atoi(v)
	return c
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	tinput := time.Now()

	b, err := ioutil.ReadFile("d10.data")
	check(err)
	data := strings.Split(string(b), "\n")
	n := []int{}
	dev := 0
	for _, f := range data {
		d := mAtoi(f)
		if d > dev {
			dev = d
		}
		n = append(n, d)
	}
	sort.Ints(n)
	dev += 3
	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)
	tinput = time.Now()
	one := 1
	three := 1
	for d := range n {

		if d+1 >= len(n) {
			break
		}
		if n[d+1]-n[d] == 1 {
			one++
		}
		if n[d+1]-n[d] == 3 {
			three++
		}
	}

	c := one * three
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println(c)
	tinput = time.Now()
	n = append(n, dev)
	//sort.Sort(sort.Reverse(sort.IntSlice(n)))
	t := map[int]int{}
	t[0] = 1
	for _, v := range n {
		t[v] = t[v-1] + t[v-2] + t[v-3]
	}
	c = t[n[len(n)-1]]
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println(c)
}
