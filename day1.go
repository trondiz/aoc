package main

import (
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

func main() {
	b, err := ioutil.ReadFile("d1.data")
	check(err)
	expense := strings.Split(string(b), "\n")

	t_input := time.Now()
out:
	for _, f := range expense {
		if f != "" {
			for _, c := range expense {
				if f != "" {
					if mAtoi(c)+mAtoi(f) == 2020 {
						input_t_elapsed := time.Since(t_input)
						log.Println("Time:", input_t_elapsed)
						log.Println(mAtoi(c) * mAtoi(f))
						break out
					}
				}
			}
		}
	}
	t_input = time.Now()
out2:
	for _, f := range expense {
		if f != "" {
			for _, c := range expense {
				for _, d := range expense {
					if f != "" {
						if mAtoi(c)+mAtoi(d)+mAtoi(f) == 2020 {
							input_t_elapsed := time.Since(t_input)
							log.Println("Time:", input_t_elapsed)
							log.Println(mAtoi(c) * mAtoi(f) * mAtoi(d))
							break out2
						}
					}
				}
			}
		}
	}
	input_t_elapsed := time.Since(t_input)
	log.Println("Time:", input_t_elapsed)
}
