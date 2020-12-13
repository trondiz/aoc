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

	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)
	min := 9999999999
	busmin := 0
	for _, b := range buses {
		if b == 0 {
			continue
		}
		nmin := ((timestamp / b) + 1) * b
		if nmin < min {
			min = nmin
			busmin = b
		}
	}
	c := busmin * (min - timestamp)

	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println(c)
	tinput = time.Now()

	bust := map[int]int{}

	for n, b := range buses {
		if b == 0 {
			continue
		}
		bust[n] = b
	}
	earliest := bust[0]
	jump := bust[0]
	for dt, bus := range bust {
		for {
			if (earliest+dt)%bus == 0 {
				break
			}
			earliest += jump
		}
		jump = findJump(jump, bus)

	}

	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println(earliest)
}

func findJump(a, b int) int {
	c := a
	d := b
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	f := a
	return c * d / f
}
