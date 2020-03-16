package main

func main() {
	var catalogueLines = getCatalogueEntries("./testCatalogue.txt")
	for _, v := range catalogueLines {
		ctLine(v)
	}
}
