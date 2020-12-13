package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"strconv"
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
func miToa(v int) string {
	c := strconv.Itoa(v)
	return c
}

type pos struct {
	x, y int
}

func main() {
	tinput := time.Now()

	inputtelapsed := time.Since(tinput)
	log.Println("Read data:", inputtelapsed)
	tinput = time.Now()

	c := 0
	s := ""
	for {
		c++
		s = "yzbqklnj" + miToa(c)
		f := md5.Sum([]byte(s))
		sm := hex.EncodeToString(f[:])
		if strings.HasPrefix(sm, "00000") {
			fmt.Println(sm)
			break
		}
	}

	inputtelapsed = time.Since(tinput)
	log.Println("Time p1:", inputtelapsed)
	log.Println(c)
	tinput = time.Now()

	c = 0
	s = ""
	for {
		c++
		s = "yzbqklnj" + miToa(c)
		f := md5.Sum([]byte(s))
		sm := hex.EncodeToString(f[:])
		if strings.HasPrefix(sm, "000000") {
			fmt.Println(sm)
			break
		}
	}
	inputtelapsed = time.Since(tinput)
	log.Println("Time p2:", inputtelapsed)
	log.Println(c)
}
