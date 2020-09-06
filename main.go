package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "test"
)

var (
	db *sql.DB
)

func main() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
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
	http.HandleFunc("/getCart", getCart)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getCart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: read cart")
	reqBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		handlerError(w, error)
	} 
	var cart int
	json.Unmarshal(reqBody, &cart)
	fmt.Println("id: ", cart)
	completeCart, err := getAllCart(db, cart)
	if err != nil {
		handlerError(w, err)
	} else {
		fmt.Fprintf(w, "%+v", completeCart)
	}

}

func listItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: list items")
	items, err := listDbItems(db)
	if err != nil {
		handlerError(w, err)
	} else {
		fmt.Fprintf(w, "%+v", items)
	}

}

func createcart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create Cart!")
	reqBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		handlerError(w, error)
	} else {
		var createdCart Cart
		json.Unmarshal(reqBody, &createdCart)
		newCart, err := createCartDb(createdCart, db)
		if err != nil {
			handlerError(w, err)
		}
		fmt.Fprintf(w, "%+v", newCart)
		fmt.Println("Endpoint Hit: create cart")
	}

}

func addItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "add Item!")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handlerError(w, err)
	} else {
		var newItem CreateItem
		err = json.Unmarshal(reqBody, &newItem)
		if err != nil {
			fmt.Println("Error Unmarshal json")
		}
		fmt.Printf("%+v", newItem)

		response, errInsert := addItemToCart(newItem, db)
		if errInsert != nil {
			handlerError(w, errInsert)
		}

		fmt.Fprintf(w, "%+v", response)
		fmt.Println("Endpoint Hit: add item")
	}

}

func modifyItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "modify Item!")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handlerError(w, err)
	} else {
		var newItem CreateItem
		err = json.Unmarshal(reqBody, &newItem)
		if err != nil {
			fmt.Println("Error Unmarshal json")
		}
		fmt.Printf("%+v", newItem)

		response, errInsert := modifyItem(newItem, db)
		if errInsert != nil {
			handlerError(w, errInsert)
		}

		fmt.Fprintf(w, "%+v", response)
		fmt.Println("Endpoint Hit: modify item")
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
	fmt.Fprintf(w, "%+v", err)
}
