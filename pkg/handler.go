package shoppingpal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func ReturnAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllItems")
	source := Source{
		Name:          "N1",
		Author:        "A1",
		Title:         "T1",
		Content:       "C1",
		URL:           "https://test",
		DatePublished: "09-08-2019",
	}

	item1 := Item{
		Name:    "Book",
		Price:   40,
		Summary: "My Book",
		Sources: source,
	}

	item2 := Item{
		Name:    "Book2",
		Price:   20,
		Summary: "My Book 2",
		Sources: source,
	}

	ItemList := Items{
		item1,
		item2,
	}
	json.NewEncoder(w).Encode(ItemList)
}
