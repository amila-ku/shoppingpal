package shoppingpal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/amila-ku/shoppingpal/api/docs"
	"github.com/gorilla/mux"
	swagger "github.com/swaggo/http-swagger"
)

var ItemList = NewItems()

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
	myRouter.HandleFunc("/swagger", swaggerdocs)
	myRouter.HandleFunc("/items", returnAllItems).Methods("GET")
	myRouter.HandleFunc("/item/{id}", returnSingleItem).Methods("GET")
	myRouter.HandleFunc("/item/{id}", deleteItem).Methods("DELETE")
	myRouter.HandleFunc("/item", createNewItem).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func swaggerdocs() {

	swagger.Handler(
		swagger.URL("http://localhost:10000/swagger/doc.json"), //The url pointing to API definition"
	)

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
	//ItemList.append(itm)
	ItemList = append(ItemList, itm)

	// save to db
	//db := database{"eu-central-1", "test", "http://localhost:8000"}
	db, err := NewTable("itemtable")

	if err != nil {
		log.Fatal("Failed to create table", err)
	}
	err = db.createItem(itm)
	if err != nil {
		log.Fatal("Unable to insert item", err)
	}

	// dynamoTable, err := newDynamoTable("itemTable", "")
	// if err != nil {
	// 	log.Fatal("Unable to create table", err)
	// }

	// err = dynamoTable.Put(itm).Run()

	// if err != nil {
	// 	log.Fatal("Unable to insert item", err)
	// }

	fmt.Println(ItemList)

	prettyJSON(w, itm)

}

func returnSingleItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	//Check items slice for matching item
	for _, item := range ItemList {

		if item.Id == key {
			prettyJSON(w, item)
		}
	}
}

func returnAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllItems")

	//json.NewEncoder(w).Encode(ItemList)

	// Print Json with indents, the pretty way:
	prettyJSON(w, ItemList)

}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// parse the path parameters
	vars := mux.Vars(r)
	// extract the `id` of the item
	id := vars["id"]

	//loop through all our items
	for index, item := range ItemList {
		// delete if item id matches
		if item.Id == id {
			ItemList = append(ItemList[:index], ItemList[index+1:]...)
		}
	}

}

func prettyJSON(w http.ResponseWriter, list interface{}) {
	pretty, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(pretty)
}
