package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// loglevel 0: no logging
// loglevel 1: logging on standardised format <result><host><url>
// loglevel > 1: Verbose human logging with response strings.

var loglevel = 1

// ComRoute is the struct holding host-url combinations that shall be tested for contact.
type ComRoute struct {
	host string
	url  string
}

func getHost() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}

func getCatalogueEntries(cataloguePath string) []string {
	file, err := os.Open(cataloguePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return (lines)
}

func doReq(cr *ComRoute) bool {
	if cr.host == getHost() {
		resp, err := http.Get(cr.url)
		if err != nil {
			if loglevel > 1 {
				fmt.Printf("\nContact call to %s FAILED. %s is supposed to be able to call this endpoint", cr.url, cr.host)
			}
			if loglevel == 1 {
				fmt.Printf("\n<FAIL><%s><%s>", cr.host, cr.url)
			}
			return false
		}

		if loglevel > 1 {
			fmt.Println(resp)
		}
		if loglevel == 1 {
			fmt.Printf("\n<PASS><%s><%s>", cr.host, cr.url)
		}
	} else {
		if loglevel > 1 {
			fmt.Printf("\nNot doing contact call to %s. %s is not supposed to be able to call this endpoint", cr.url, cr.host)
		}
	}
	return true
}

func ct(host string, url string) {
	c := new(ComRoute)
	c.host = host
	c.url = url
	doReq(c)
}

func ctLine(hostAndURLLine string) {
	splitted := strings.SplitN(hostAndURLLine, " ", 2)
	ct(splitted[0], splitted[1])
}
