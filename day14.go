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

// Sets the bit at pos in the integer n.
func setBit(n int, pos int) int {
	n |= (1 << pos)
	return n
}

// Clears the bit at pos in n.
func clearBit(n int, pos int) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	tinput := time.Now()
	k, err := ioutil.ReadFile("d14.data")
	check(err)
	data := strings.Split(string(k), "\n")

	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)

	mask := ""

	memory := map[int]int{}
	for _, d := range data {
		if strings.HasPrefix(d, "mask") {
			fmt.Sscanf(d, "mask = %s", &mask)
			mask = reverse(mask)
		}

		if strings.HasPrefix(d, "mem") {
			addr := 0
			val := 0
			fmt.Sscanf(d, "mem[%d] = %d", &addr, &val)
			//fmt.Println(mask)
			//fmt.Printf("%036b", val)
			//fmt.Println()
			for i := 35; i >= 0; i-- {
				switch mask[i] {
				case '0':
					//fmt.Println("Flipping: ", i, " in ", val)
					val = clearBit(val, i)
				case '1':
					//fmt.Println("Flipping: ", i, " in ", val)
					val = setBit(val, i)
				}
			}
			//fmt.Printf("%036b", val)
			//fmt.Println()

			memory[addr] = val
		}
	}
	c := 0
	for _, val := range memory {
		c += val
	}
	//bitmask = [36]bit

	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println(c)
	tinput = time.Now()

	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println(c)
}
