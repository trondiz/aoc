package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	tinput := time.Now()

	b, err := ioutil.ReadFile("d6.data")
	check(err)
	groups := strings.Split(string(b), "\n\n")

	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)
	tinput = time.Now()
	c := 0

	for _, v := range groups {
		arb := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
		for _, q := range v {
			if q == 10 {
				continue
			}
			arb[q-97] = true
		}
		for _, ar := range arb {
			if ar {
				c++
			}
		}
	}

	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println(c)
	tinput = time.Now()
	c = 0

	for _, v := range groups {
		ln := 1
		arb := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for _, q := range v {
			if q == 10 {
				ln++
				continue
			}
			arb[q-97]++
		}
		for _, ar := range arb {
			if ar == ln {
				c++
			}
		}
	}

	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println(c)
}
