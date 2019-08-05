package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/amila-ku/shoppingpal"
)


func main() {

	source := Source{
		Name: "N1",
		Author: "A1",
		Title: "T1",
		Content: "C1",
		URL: "https://test",
		DatePublished: "09-08-2019"

	}

	item1 := Item{
		Name: "Book",
		Price: "$20",
		Title: "My Book"

	}

	item2 := Item{
		Name: "Book2",
		Price: "$20",
		Title: "My Book 2"

	}

	ItemList := []ItemList {
		item1,
		item2,
	}

	handleRequests()
}
