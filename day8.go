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

type instr struct {
	cmd  string
	seen bool
	val  int
}

func runprog(ins []instr) int {
	pos := 0
	acc := 0
	for ins[pos].seen == false {
		ins[pos].seen = true
		switch ins[pos].cmd {
		case "acc":
			acc += ins[pos].val
			pos++
		case "jmp":
			pos += ins[pos].val
		default:
			pos++
		}
		if pos >= len(ins) {
			return acc
		}
	}
	return 0
}

func main() {
	tinput := time.Now()

	b, err := ioutil.ReadFile("d8.data")
	check(err)
	data := strings.Split(string(b), "\n")
	ins := []instr{}
	for _, f := range data {
		s := ""
		d := 0
		fmt.Sscanf(f, "%s %d", &s, &d)

		ins = append(ins, instr{cmd: s, val: d, seen: false})
	}
	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)
	tinput = time.Now()
	acc := 0
	pos := 0
	for ins[pos].seen == false {
		ins[pos].seen = true
		switch ins[pos].cmd {
		case "acc":
			acc += ins[pos].val
			pos++
		case "jmp":
			pos += ins[pos].val
		default:
			pos++
		}
	}
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println(acc)
	tinput = time.Now()
	for i := range ins {
		acc = 0
		if ins[i].cmd == "acc" {
			continue
		}
		for is := range ins {
			ins[is].seen = false
		}
		if ins[i].cmd == "jmp" {
			ins[i].cmd = "nop"
		} else {
			ins[i].cmd = "jmp"
		}
		acc = runprog(ins)
		if acc > 0 {
			break
		}
		if ins[i].cmd == "jmp" {
			ins[i].cmd = "nop"
		} else {
			ins[i].cmd = "jmp"
		}
	}

	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println(acc)
}
