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

func mAtoia(v []string) []int {
	l := make([]int, len(v))
	for i, s := range v {
		l[i], _ = strconv.Atoi(s)
	}
	return l
}

type rule struct {
}

func main() {
	tinput := time.Now()

	k, err := ioutil.ReadFile("d16.data")
	check(err)
	data := strings.Split(string(k), "\n")
	rulesName := []string{}
	rules := [][]int{}
	tickets := [][]int{}
	validTickets := [][]int{}
	readmode := "rules"
	myTicket := []int{}
	for i, f := range data {
		switch readmode {
		case "rules":
			if f == "" {
				readmode = "searchtickets"
				continue
			}
			re := pcre.MustCompile(`^([A-Za-z\s]+): (\d+)-(\d+) or (\d+)-(\d+).*`, 0).MatcherString(f, 0)
			ranges := []int{}
			for i := 2; i <= re.Groups(); i++ {
				ranges = append(ranges, mAtoi(re.GroupString(i)))
			}
			rulesName = append(rulesName, re.GroupString(1))
			rules = append(rules, ranges)
		case "searchtickets":
			if f == "your ticket:" {
				myTicket = mAtoia(strings.Split(data[i+1], ","))
			}
			if f != "nearby tickets:" {
				continue
			} else {
				readmode = "parsetickets"
			}
		case "parsetickets":
			tickets = append(tickets, (mAtoia(strings.Split(f, ","))))
		}
	}

	inputtelapsed := time.Since(tinput)
	fmt.Println("Read data:", inputtelapsed)
	tinput = time.Now()
	c := 0
	for _, v := range tickets {
		for _, vv := range v {
			valid := false
			for _, r := range rules {
				if vv >= r[0] && vv <= r[1] {
					valid = true
					break
				}
				if vv >= r[2] && vv <= r[3] {
					valid = true
					break
				}
			}
			if !valid {
				c += vv
			}
		}
	}

	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p1:", inputtelapsed)
	fmt.Println(c)
	tinput = time.Now()
	c = 1
	for _, v := range tickets {
		validTicket := true
		for _, vv := range v {
			valid := false
			for _, r := range rules {
				if vv >= r[0] && vv <= r[1] {
					valid = true
					break
				}
				if vv >= r[2] && vv <= r[3] {
					valid = true
					break
				}
			}
			if !valid {
				validTicket = false
			}
		}
		if validTicket {
			validTickets = append(validTickets, v)
		}
	}
	rulem := [][]int{}
	for i, r := range rules {
		rulem = append(rulem, []int{})
		for pos := 0; pos < len(myTicket); pos++ {
			posmatchr := true
			for _, t := range validTickets {
				valid := false
				if t[pos] >= r[0] && t[pos] <= r[1] {
					valid = true
				}
				if t[pos] >= r[2] && t[pos] <= r[3] {
					valid = true
				}
				if !valid {
					posmatchr = false
					break
				}
			}
			if posmatchr {
				rulem[i] = append(rulem[i], pos)
			}
		}
	}
	rulemap := make([]int, len(rules))
	for range rulemap {
		for j, r := range rulem {
			if len(r) == 1 {
				known := r[0]
				rulemap[j] = known
				for i, ri := range rulem {
					nr := []int{}
					for m := range ri {
						if ri[m] != known {
							nr = append(nr, ri[m])
						}
					}
					rulem[i] = nr
				}
			}
		}
	}
	for ii, i := range rulemap {
		if strings.HasPrefix(rulesName[ii], "departure") {
			c *= myTicket[i]
		}
	}
	inputtelapsed = time.Since(tinput)
	fmt.Println("Time p2:", inputtelapsed)
	fmt.Println(c)
}
