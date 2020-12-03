package main

import (
	"io/ioutil"
	"log"

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

func passIsValid(min int, max int, letter string, pass string) bool {
	c := strings.Count(pass, letter)
	if c <= max && c >= min {
		return true
	}
	return false
}

func passIsValid2(min int, max int, letter rune, pass []rune) bool {
	if (pass[min] == letter) != (pass[max] == letter) {
		return true
	}
	return false
}

type p1 struct {
	min    int
	max    int
	letter string
	pass   string
}
type p2 struct {
	min    int
	max    int
	letter rune
	pass   []rune
}

func main() {
	tinput := time.Now()

	b, err := ioutil.ReadFile("d2.data")
	check(err)
	expense := strings.Split(string(b), "\n")
	t1 := []p1{}
	t2 := []p2{}
	for _, f := range expense {
		// 9-13 j: jjjjjjjjjjjjf
		re := pcre.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<letter>.): (?P<pass>.+)$`, 0)
		res := re.MatcherString(f, 0)
		t1 = append(t1, p1{mAtoi(res.GroupString(1)) - 1, mAtoi(res.GroupString(2)) - 1, res.GroupString(3), res.GroupString(4)})
		t2 = append(t2, p2{mAtoi(res.GroupString(1)) - 1, mAtoi(res.GroupString(2)) - 1, []rune(res.GroupString(3))[0], []rune(res.GroupString(4))})
	}

	inputtelapsed := time.Since(tinput)
	log.Println("Read data:", inputtelapsed)
	tinput = time.Now()
	c := 0
	for _, r := range t1 {
		if passIsValid(r.min, r.max, r.letter, r.pass) {
			c++
		}
	}
	inputtelapsed = time.Since(tinput)
	log.Println("Time p1:", inputtelapsed)
	log.Println(c)

	tinput = time.Now()
	c = 0
	for _, r := range t2 {
		if passIsValid2(r.min, r.max, r.letter, r.pass) {
			c++
		}
	}
	inputtelapsed = time.Since(tinput)
	log.Println("Time p2:", inputtelapsed)
	log.Println(c)
}
