package main

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve"
)

// START OMIT
func main() {

	index, err := bleve.Open("gophercon.bleve")
	if err != nil {
		log.Fatal(err)
	}

	q := bleve.NewMatchAllQuery()
	req := bleve.NewSearchRequest(q)
	req.Size = 0
	req.AddFacet("sections", // HL
		bleve.NewFacetRequest("section", 50)) // HL
	res, err := index.Search(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

// END OMIT
