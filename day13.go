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

func main() {
	tinput := time.Now()
	k, err := ioutil.ReadFile("d13.data")
	check(err)
	data := strings.Split(string(k), "\n")
	buses := []int{}
	timestamp := 0
	for i, f := range data {
		if i == 0 {
			timestamp = mAtoi(f)
		} else {
			busestmp := strings.Split(string(f), ",")
			for _, b := range busestmp {
				buses = append(buses, mAtoi(b))
			}
		}
	}
	fmt.Println(buses)

	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)

	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println()
	tinput = time.Now()

	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println()
}
