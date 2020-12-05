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

func valid(p string) bool {
	c := strings.Count(p, ":")
	if c == 8 {
		return true
	}
	if c == 7 {
		if !strings.Contains(p, "cid") {
			return true
		}
	}
	return false
}

func valid2(p string) bool {
	if valid(p) {

		byr := mAtoi(pcre.MustCompile(`byr:(\d+)`, 0).MatcherString(p, 0).GroupString(1))
		if byr > 2002 || byr < 1920 {
			return false
		}
		iyr := mAtoi(pcre.MustCompile(`iyr:(\d+)`, 0).MatcherString(p, 0).GroupString(1))
		if iyr > 2020 || iyr < 2010 {
			return false
		}
		eyr := mAtoi(pcre.MustCompile(`eyr:(\d+)`, 0).MatcherString(p, 0).GroupString(1))
		if eyr > 2030 || eyr < 2020 {
			return false
		}
		hgtcm := mAtoi(pcre.MustCompile(`hgt:(\d+)cm`, 0).MatcherString(p, 0).GroupString(1))
		hgtin := 0
		if hgtcm == 0 {
			hgtin = mAtoi(pcre.MustCompile(`hgt:(\d+)in`, 0).MatcherString(p, 0).GroupString(1))
		}
		if !((hgtcm >= 150 && hgtcm <= 193) || (hgtin >= 59 && hgtin <= 76)) {
			return false
		}
		hcl := pcre.MustCompile(`hcl:#[0-9a-f]{6}`, 0).MatcherString(p, 0).Matches()
		if !hcl {
			return false
		}
		ecl := pcre.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)`, 0).MatcherString(p, 0).Matches()
		if !ecl {
			return false
		}
		pid := pcre.MustCompile(`pid:\d{9}(\s+|$)`, 0).MatcherString(p, 0).Matches()
		if !pid {
			return false
		}
		return true
	}

	return false
}

func main() {
	tinput := time.Now()

	b, err := ioutil.ReadFile("d4.data")
	check(err)
	psprts := strings.Split(string(b), "\n\n")

	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)
	tinput = time.Now()
	c := 0
	for _, y := range psprts {
		if valid(y) {
			c++
		}

	}
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println(c)

	tinput = time.Now()
	c = 0
	for _, y := range psprts {
		if valid2(y) {
			c++
		}
	}
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println(c)
}
