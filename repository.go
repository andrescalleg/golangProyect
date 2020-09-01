package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "test"
)

func getCart() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	sqlStatement :=
		"select cart_id, user_name " +
			"from public.cart " +
			"where cart_id = ?; "
	sqlStatement = strings.Replace(sqlStatement, "?", "1", 1)
	fmt.Println(sqlStatement)

	var cart Cart
	row := db.QueryRow(sqlStatement)
	switch err := row.Scan(&cart.id, &cart.UserName); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(cart)
	default:
		panic(err)
	}

}

func getItem() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	sqlStatement :=
		"select item_id, item_name, item_value " +
			"from public.item " +
			"where item_id = ?; "
	sqlStatement = strings.Replace(sqlStatement, "?", "1", 1)
	fmt.Println(sqlStatement)

	var item Item
	row := db.QueryRow(sqlStatement)
	switch err := row.Scan(&item.id, &item.name, &item.value); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(item)
	default:
		panic(err)
	}
}

func getAllCart() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	sqlStatement :=
		"SELECT cart.cart_id, cart.user_name, cart_items.item_amount, item.item_id, item.item_name, item.item_value " +
			"FROM cart_items, cart, item " +
			"WHERE cart.cart_id = ? AND cart_items.cart_id = cart.cart_id AND item.item_id = cart_items.item_id; "
	sqlStatement = strings.Replace(sqlStatement, "?", "1", 1)
	fmt.Println(sqlStatement)

	var item Item
	var cart Cart
	row := db.QueryRow(sqlStatement)
	switch err := row.Scan(&cart.id, &cart.UserName, &item.amount, &item.id, &item.name, &item.value); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		cart.Products = append(cart.Products, item)
		fmt.Println(cart)
	default:
		panic(err)
	}

}
