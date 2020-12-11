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

func Equal(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	tinput := time.Now()

	k, err := ioutil.ReadFile("d11.data")
	check(err)
	data := strings.Split(string(k), "\n")
	b := string(k)
	width := len(data[0]) + 1

	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)
	a := 0
	ops := map[int][]rune{}
	pass := 0
out:
	for {
		if len(ops) != 0 {
			if Equal(ops[pass-2], ops[pass-1]) {
				n := 0
				for g := range ops[pass-1] {
					if ops[pass-1][g] == '#' {
						n++
					}
				}
				a = n
				break out
			}
		}
		for f := range b {
			c := 0
			if f-width-1 >= 0 && b[f-width-1] == '#' {
				c++
			}
			if f-width >= 0 && b[f-width] == '#' {
				c++
			}
			if f-width+1 >= 0 && b[f-width+1] == '#' {
				c++
			}
			if f-1 >= 0 && b[f-1] == '#' {
				c++
			}
			if f+1 < len(b) && b[f+1] == '#' {
				c++
			}
			if f+width-1 < len(b) && b[f+width-1] == '#' {
				c++
			}
			if f+width < len(b) && b[f+width] == '#' {
				c++
			}
			if f+width+1 < len(b) && b[f+width+1] == '#' {
				c++
			}
			if c == 0 && b[f] == 'L' {
				ops[pass] = append(ops[pass], '#')
			} else if c >= 4 && b[f] == '#' {
				ops[pass] = append(ops[pass], 'L')
			} else {
				ops[pass] = append(ops[pass], rune(b[f]))
			}
		}
		b = string(ops[pass])
		pass++
	}
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println(a)
	tinput = time.Now()
	b = string(k)
	a = 0
	ops = map[int][]rune{}
	pass = 0
out2:
	for {
		if len(ops) != 0 {
			if Equal(ops[pass-2], ops[pass-1]) {
				n := 0
				for g := range ops[pass-1] {
					if ops[pass-1][g] == '#' {
						n++
					}
				}
				a = n
				break out2
			}
		}
		for f := range b {
			c := 0
			for i := 1; f-(width-1)*i >= 0; i++ {
				if b[f-(width-1)*i] == '#' {
					c++
					break
				}
				if b[f-(width-1)*i] == 'L' {
					break
				}
				if b[f-(width-1)*i] == 10 {
					break
				}
			}
			for i := 1; f-(width*i) >= 0; i++ {
				if b[f-(width*i)] == '#' {
					c++
					break
				}
				if b[f-(width*i)] == 'L' {
					break
				}
				if b[f-(width*i)] == 10 {
					break
				}
			}
			for i := 1; f-(width+1)*i >= 0; i++ {
				if b[f-(width+1)*i] == '#' {
					c++
					break
				}
				if b[f-(width+1)*i] == 'L' {
					break
				}
				if b[f-(width+1)*i] == 10 {
					break
				}
			}
			for i := 1; f-i >= 0; i++ {
				if b[f-i] == '#' {
					c++
					break
				}
				if b[f-i] == 'L' {
					break
				}
				if b[f-i] == 10 {
					break
				}
			}
			for i := 1; f+i < len(b); i++ {
				if b[f+i] == '#' {
					c++
					break
				}
				if b[f+i] == 'L' {
					break
				}
				if b[f+i] == 10 {
					break
				}
			}
			for i := 1; f+(width-1)*i < len(b); i++ {
				if b[f+(width-1)*i] == '#' {
					c++
					break
				}
				if b[f+(width-1)*i] == 'L' {
					break
				}
				if b[f+(width-1)*i] == 10 {
					break
				}
			}
			for i := 1; f+(width)*i < len(b); i++ {
				if b[f+(width)*i] == '#' {
					c++
					break
				}
				if b[f+(width)*i] == 'L' {
					break
				}
				if b[f+(width)*i] == 10 {
					break
				}
			}
			for i := 1; f+(width+1)*i < len(b); i++ {
				if b[f+(width+1)*i] == '#' {
					c++
					break
				}
				if b[f+(width+1)*i] == 'L' {
					break
				}
				if b[f+(width+1)*i] == 10 {
					break
				}
			}
			if c == 0 && b[f] == 'L' {
				ops[pass] = append(ops[pass], '#')
			} else if c >= 5 && b[f] == '#' {
				ops[pass] = append(ops[pass], 'L')
			} else {
				ops[pass] = append(ops[pass], rune(b[f]))
			}
		}
		b = string(ops[pass])
		pass++
	}

	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println(a)
}
