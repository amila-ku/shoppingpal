package main

import (
	"fmt"

	shoppingpal "github.com/amila-ku/shoppingpal/pkg"
)

func main() {
	fmt.Println("starting")
	ItemList := NewItems()

	shoppingpal.HandleRequests(ItemList)

}
