package shoppingpal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var ItemList Items

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests(I Items) {
	// http.HandleFunc("/", returnAllItems)
	// log.Fatal(http.ListenAndServe(":10000", nil))

	// replaceing http.HandleFunc with myRouter.HandleFunc
	ItemList = I
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllItems).Methods("GET")
	myRouter.HandleFunc("/item/{id}", returnSingleItem).Methods("GET")
	myRouter.HandleFunc("/item", createNewItem).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func createNewItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: CreateNewItem")

	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	var itm Item
	json.Unmarshal(reqBody, &itm)
	// update our global item array to include our new item
	ItemList.append(itm)

}

func returnSingleItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, item := range ItemList {

		if item.Id == key {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func returnAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllItems")

	json.NewEncoder(w).Encode(ItemList)
}
