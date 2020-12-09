package main

import (
	"fmt"
	"io/ioutil"

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

func validate(s []int, v int) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i]+s[j] == v {
				return true
			}
		}
	}
	return false
}
func sum(i []int) int {
	s := 0
	for _, j := range i {
		s += j
	}
	return s
}
func msum(i []int) int {
	min := 9999999999999999
	max := 0
	for _, j := range i {
		if j < min {
			min = j
		}
		if j > max {
			max = j
		}
	}
	return min + max
}

var cf = 25

func main() {
	tinput := time.Now()

	b, err := ioutil.ReadFile("d9.data")
	check(err)
	data := strings.Split(string(b), "\n")
	n := []int{}
	for _, f := range data {
		n = append(n, mAtoi(f))
	}
	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)
	tinput = time.Now()
	c := 0
	a := 0
	for i := cf; i < len(n); i++ {
		if !validate(n[i-cf:i], n[i]) {
			c = n[i]
			a = i
		}
	}
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println(c)
	tinput = time.Now()
	wi := 2
	pos := 0
out:
	for {
		for i := 0; i < a; i++ {
			if i+wi > len(n) {
				break
			}
			if sum(n[i:i+wi]) == c {
				pos = i
				break out
			}
		}
		wi++
	}

	w := msum(n[pos : pos+wi])
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println(w)
}
