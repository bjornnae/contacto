package main

import (
	"fmt"
	"time"
)

func main() {
	var hostname = getHost()
	now := time.Now()
	fmt.Printf("<RUN><%s><%s>", hostname, now)
	var catalogueLines = getCatalogueEntries("./testCatalogue.txt")
	for _, v := range catalogueLines {
		ctLine(v)
	}
}
