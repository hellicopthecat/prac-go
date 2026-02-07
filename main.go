package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/hellicopthecat/learngo/blockchain"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}
	data := homeData{"homemomomo", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}

const port string = ":4000"

func main() {
	http.HandleFunc("/", home)

	fmt.Printf("Listen on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
