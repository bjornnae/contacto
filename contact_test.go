package contacto

import (
	"fmt"
	"testing"
)

func TestGetHost(t *testing.T) {
	got := getHost()
	fmt.Println("I am host:" + got)
	return
}

func TestRoute1(t *testing.T) {
	host := getHost()
	r1 := new(ComRoute)
	r1.host = host
	r1.url = "http://localhost:8080/api/hello"
	doReq(r1)
}

func TestRoute2(t *testing.T) {
	r2 := new(ComRoute)
	r2.host = "garkast3"
	r2.url = "http://localhost:8080/api/hello"
	doReq(r2)
}

// This is a failing
func TestRoute3(t *testing.T) {
	ct(getHost(), "http://localhost:8081/api/hello")
}

func TestReader(t *testing.T) {
	s := getCatalogueEntries("./testCatalogue.txt")
	for _, v := range s {
		fmt.Println(v)
	}
}
