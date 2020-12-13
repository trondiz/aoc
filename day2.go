package main

import (
	"fmt"
	"io/ioutil"
	"log"

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

type gift struct {
	l, w, h int
}

func main() {
	tinput := time.Now()

	b, err := ioutil.ReadFile("d2.data")
	check(err)
	gifts := strings.Split(string(b), "\n")
	g := []gift{}

	for _, f := range gifts {
		l := 0
		w := 0
		h := 0
		fmt.Sscanf(f, "%dx%dx%d", &l, &w, &h)
		g = append(g, gift{l: l, w: w, h: h})
	}
	fmt.Println(g)
	inputtelapsed := time.Since(tinput)
	log.Println("Read data:", inputtelapsed)
	tinput = time.Now()
	//2*l*w + 2*w*h + 2*h*l.
	c := 0
	for _, gi := range g {
		s1 := gi.l * gi.w
		s2 := gi.w * gi.h
		s3 := gi.h * gi.l
		c += (2 * s1) + (2 * s2) + (2 * s3)
		if s1 <= s2 && s1 <= s3 {
			c += s1
			continue
		}
		if s2 <= s1 && s2 <= s3 {
			c += s2
			continue
		}
		if s3 <= s2 && s3 <= s1 {
			c += s3
			continue
		}

	}
	inputtelapsed = time.Since(tinput)
	log.Println("Time p1:", inputtelapsed)
	log.Println(c)
	tinput = time.Now()

	c = 0
	for _, gi := range g {
		c += (gi.l * gi.w * gi.h)
		s1 := gi.l
		s2 := gi.w
		s3 := gi.h
		if s1 >= s2 && s1 >= s3 {
			c += (2 * s2) + (2 * s3)
			continue
		}
		if s2 >= s1 && s2 >= s3 {
			c += (2 * s1) + (2 * s3)
			continue
		}
		if s3 >= s2 && s3 >= s1 {
			c += (2 * s2) + (2 * s1)
			continue
		}

	}
	inputtelapsed = time.Since(tinput)
	log.Println("Time p2:", inputtelapsed)
	log.Println(c)
}
