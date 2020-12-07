package main

import (
	"fmt"
	"io/ioutil"

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

type bag struct {
	me    string
	bags  []string
	nbags []int
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func findShinyGold(mybag bag, data []bag) bool {
	if len(mybag.bags) == 0 {
		return false
	}
	for _, b := range mybag.bags {
		if strings.Contains(b, "shiny gold") {
			return true
		}
	}

	for _, b := range mybag.bags {
		for _, d := range data {
			if strings.Contains(d.me, b) {
				if findShinyGold(d, data) {
					return true
				}
			}
		}
	}
	return false
}

func countBags(mybag bag, data []bag) int {
	sum := sum(mybag.nbags)
	for i, b := range mybag.bags {
		for _, d := range data {
			if strings.Contains(d.me, b) {
				sum += countBags(d, data) * mybag.nbags[i]
			}
		}
	}
	return sum
}

func main() {
	tinput := time.Now()

	b, err := ioutil.ReadFile("d7.data")
	check(err)
	data := strings.Split(string(b), "\n")
	bags := []bag{}
	for _, f := range data {
		me := pcre.MustCompile(`^(\w+ \w+).*`, 0).MatcherString(f, 0).GroupString(1)
		sbags := []string{}
		nbags := []int{}
		sub := strings.Split(f, ",")
		for _, s := range sub {
			res := pcre.MustCompile(`.*(\d+) (\w+ \w+) (bag|bags).*`, 0).MatcherString(s, 0)
			bgs := res.GroupString(2)
			n := res.GroupString(1)
			sbags = append(sbags, bgs)
			nbags = append(nbags, mAtoi(n))
			if sbags[0] == "" {
				sbags = []string{}
			}
			if nbags[0] == 0 {
				nbags = []int{}
			}
		}
		bags = append(bags, bag{me: me, bags: sbags, nbags: nbags})
	}
	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)
	tinput = time.Now()
	c := 0
	for _, b := range bags {
		if findShinyGold(b, bags) {
			c++
		}
	}
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println(c)
	tinput = time.Now()
	for _, b := range bags {
		if strings.Contains(b.me, "shiny gold") {
			c = countBags(b, bags)
		}
	}
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println(c)
}
