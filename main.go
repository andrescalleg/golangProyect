package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/createcart", createcart)
	http.HandleFunc("/additem", addItem)
	http.HandleFunc("/listitems", listItems)
	http.HandleFunc("/modifyitem", modifyItems)
	http.HandleFunc("/removeitem", removeItems)
	http.HandleFunc("/removeall", removeAll)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	getCart()
	getItem()
	getAllCart()
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TESTPage!")
	fmt.Println("Endpoint Hit: homePage")
}

func createcart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create Cart!")
	reqBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		handlerError(w, error)
	} else {
		var createdCart Cart
		json.Unmarshal(reqBody, &createdCart)

		fmt.Fprintf(w, "%+v", createdCart)
		fmt.Println("Endpoint Hit: create cart")
	}

}

func addItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "add Item!")
	reqBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		handlerError(w, error)
	} else {
		var newItem Item
		json.Unmarshal(reqBody, &newItem)

		fmt.Fprintf(w, "%+v", newItem)
		fmt.Println("Endpoint Hit: add item")
	}

}

func listItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List items")
	fmt.Println("Endpoint Hit: list items")
}

func modifyItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "modify items")
	reqBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		handlerError(w, error)
	} else {
		var cartToModify Cart
		json.Unmarshal(reqBody, &cartToModify)

		fmt.Fprintf(w, "%+v", cartToModify)
		fmt.Println("Endpoint Hit: modify items")
	}
}

func removeItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "remove items")
	reqBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		handlerError(w, error)
	} else {
		var removeItem Item
		json.Unmarshal(reqBody, &removeItem)

		fmt.Fprintf(w, "%+v", removeItem)
		fmt.Println("Endpoint Hit: remove items")
	}
}

func removeAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Remove all")
	reqBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		handlerError(w, error)
	} else {
		var createdCart Cart
		json.Unmarshal(reqBody, &createdCart)

		fmt.Fprintf(w, "%+v", createdCart)
		fmt.Println("Endpoint Hit: remove all")
	}
}

func handlerError(w http.ResponseWriter, err error) {
	fmt.Println("Error parsing", err)
	fmt.Fprintf(w, "%+v", "bad request")
}
