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

	qString := `+content:debugger ` // HLQUERY
	qString += `title:"delve" `     // HLQUERY
	qString += `title:go~2 `        // HLQUERY
	qString += `-content:rust `     // HLQUERY
	qString += `word_count:>30`     // HLQUERY
	q := bleve.NewQueryStringQuery(qString)
	req := bleve.NewSearchRequest(q)
	req.Highlight = bleve.NewHighlightWithStyle("html")
	req.Fields = []string{"title", "author", "content", "word_count"}
	res, err := index.Search(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

// END OMIT
