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
	r := strings.NewReplacer("B", "1", "F", "0", "L", "0", "R", "1").Replace(y)
	/*r := ""
	for i := range y {
		switch y[i] {
		case 'B':
			r += "1"
		case 'F':
			r += "0"
		case 'L':
			r += "0"
		case 'R':
			r += "1"
		}
	}*/

	ri, _ := strconv.ParseInt(r, 2, 11)
	return int(ri)
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
