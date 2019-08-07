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

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllItems")
	item1 := Item{
		Name:  "Book",
		Price: "$20",
		Title: "My Book",
	}

	item2 := Item{
		Name:  "Book2",
		Price: "$20",
		Title: "My Book 2",
	}

	ItemList := []Items{
		item1,
		item2,
	}
	json.NewEncoder(w).Encode(ItemList)
}
