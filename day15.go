package main

import (
	"fmt"
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

type number struct {
	val  int
	turn int
}

func main() {
	tinput := time.Now()
	turn := 0
	last := 0
	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)

	data := []int{2, 0, 6, 12, 1, 3}
	nums := make([]int, 100000000)
	for i := 1; i < len(data); i++ {
		nums[data[i-1]] = i
	}
	last = data[len(data)-1:][0]
	turn = len(data) - 1
	current := 0
	for j := len(data); j < 2020; j++ {
		turn++
		if nums[last] == 0 {
			current = 0
		} else {
			current = turn - nums[last]
		}
		nums[last] = turn
		last = current
	}

	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println(last)
	tinput = time.Now()

	nums = make([]int, 100000000)
	for i := 1; i < len(data); i++ {
		nums[data[i-1]] = i
	}
	last = data[len(data)-1:][0]
	turn = len(data) - 1
	current = 0
	for j := len(data); j < 30000000; j++ {
		turn++
		if nums[last] == 0 {
			current = 0
		} else {
			current = turn - nums[last]
		}
		nums[last] = turn
		last = current
	}

	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println(last)
}
