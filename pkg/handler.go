package shoppingpal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests() {
	// http.HandleFunc("/", returnAllItems)
	// log.Fatal(http.ListenAndServe(":10000", nil))

	// replaceing http.HandleFunc with myRouter.HandleFunc
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllItems)
	myRouter.HandleFunc("/item/{id}", returnSingleItem)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnSingleItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)

	ItemList := GetItems()

	for _, item := range ItemList {
		if item.Id == key {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func returnAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllItems")

	ItemList := GetItems()

	json.NewEncoder(w).Encode(ItemList)
}
