package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"

	"strings"
	"time"

	pcre "github.com/gijsbers/go-pcre"
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

type action struct {
	direction string
	value     int
}

func main() {
	tinput := time.Now()

	k, err := ioutil.ReadFile("d12.data")
	check(err)
	data := strings.Split(string(k), "\n")
	actions := []action{}
	for _, f := range data {
		re := pcre.MustCompile(`^(.)(\d+)`, 0).MatcherString(f, 0)
		s := re.GroupString(1)
		d := mAtoi(re.GroupString(2))
		actions = append(actions, action{direction: s, value: d})
	}

	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)

	//state
	x := 0
	y := 0
	d := 90
	for _, a := range actions {
		switch a.direction {
		case "N":
			y += a.value
		case "S":
			y -= a.value
		case "E":
			x += a.value
		case "W":
			x -= a.value
		case "L":
			d = (((d - a.value) % 360) + 360) % 360
		case "R":
			d = (((d + a.value) % 360) + 360) % 360
		case "F":
			switch d {
			case 0:
				y += a.value
			case 90:
				x += a.value
			case 180:
				y -= a.value
			case 270:
				x -= a.value
			}
		}
	}
	c := math.Abs(float64(x)) + math.Abs(float64(y))
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println(c)
	tinput = time.Now()

	//state
	wpx := 10
	wpy := 1
	sx := 0
	sy := 0
	for _, a := range actions {
		switch a.direction {
		case "N":
			wpy += a.value
		case "S":
			wpy -= a.value
		case "E":
			wpx += a.value
		case "W":
			wpx -= a.value
		case "L":
			switch a.value {
			case 90:
				ny := wpx
				nx := -wpy
				wpx = nx
				wpy = ny
			case 180:
				ny := -wpy
				nx := -wpx
				wpx = nx
				wpy = ny
			case 270:
				ny := -wpx
				nx := wpy
				wpx = nx
				wpy = ny
			}
		case "R":
			switch a.value {
			case 90:
				ny := -wpx
				nx := wpy
				wpx = nx
				wpy = ny
			case 180:
				ny := -wpy
				nx := -wpx
				wpx = nx
				wpy = ny
			case 270:
				ny := wpx
				nx := -wpy
				wpx = nx
				wpy = ny
			}
		case "F":
			sy += a.value * wpy
			sx += a.value * wpx
		}
	}
	c = math.Abs(float64(sx)) + math.Abs(float64(sy))
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println(c)
}
