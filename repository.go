package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

func getItem(db *sql.DB) {
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

func getAllCart(db *sql.DB) {
	sqlStatement :=
		"SELECT cart.cart_id, cart.user_name, cart_items.item_amount, item.item_id, item.item_name, item.item_value " +
			"FROM cart_items, cart, item " +
			"WHERE cart.cart_id = ? AND cart_items.cart_id = cart.cart_id AND item.item_id = cart_items.item_id; "
	sqlStatement = strings.Replace(sqlStatement, "?", "1", 1)
	fmt.Println(sqlStatement)

	var item Item
	var cart Cart
	row := db.QueryRow(sqlStatement)
	switch err := row.Scan(&cart.id, &cart.UserName, &item.id, &item.name, &item.value); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		cart.Products = append(cart.Products, item)
		fmt.Println(cart)
	default:
		panic(err)
	}

}

func listDbItems(db *sql.DB) ([]Item, error) {

	var specificError error
	sqlStatement := fmt.Sprintf(`
		select item_id, item_name, item_value 
		from public.item`)

	fmt.Println(sqlStatement)

	var item []Item
	row, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		specificError = err
	} else {
		for row.Next() {
			var newItem Item
			err := row.Scan(&newItem.id, &newItem.name, &newItem.value)
			if err != nil {
				fmt.Println(err)
				specificError = err
			}
			item = append(item, newItem)
		}
	}

	fmt.Println(item)
	return item, specificError

}

func createCartDb(createCart Cart, db *sql.DB) (Cart, error) {
	var specificError error
	sqlStatement := fmt.Sprintf(`
	INSERT INTO public.cart(
		cart_id, user_name)
		VALUES (%d, '%s');
		`, createCart.id, createCart.UserName)
	fmt.Println(sqlStatement)

	var cart Cart
	result, err := db.Exec(sqlStatement)
	_, errInsert := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		specificError = err
	}
	if errInsert != nil {
		fmt.Println(errInsert)
		specificError = errInsert
	}
	return cart, specificError
}

func findItem(newItem CreateItem, db *sql.DB) error {
	sqlStatement := fmt.Sprintf(`
		select item_id, item_name, item_value 
		from public.item 
		where item_id = ?;`, newItem.id)
	fmt.Println(sqlStatement)

	var item Item
	row := db.QueryRow(sqlStatement)
	err := row.Scan(&item.id, &item.name, &item.value)
	return err

}

func addItemToCart(newItem CreateItem, db *sql.DB) (Item, error) {
	var specificError error
	err := findItem(newItem, db)
	if err != nil {
		fmt.Println(err)
		specificError = err
	}
	sqlStatement := fmt.Sprintf(`
	INSERT INTO public.item(
		item_id, item_name, item_value)
		VALUES (?, ?, ?);
		`, newItem.id, newItem.name, newItem.value)
	fmt.Println(sqlStatement)

	result, err := db.Exec(sqlStatement)
	_, affected := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		specificError = err
	}
	if affected != nil {
		fmt.Println(affected)
		specificError = affected
	}

	sqlStatementInsert := fmt.Sprintf(`
	INSERT INTO public.cart_items(
		item_id, cart_id, item_amount)
		VALUES (?, ?, ?);
		`, newItem.id, newItem.cartId, newItem.amount)
	fmt.Println(sqlStatementInsert)

	var item Item
	result, err = db.Exec(sqlStatementInsert)
	_, affected = result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		specificError = err
	}
	if affected != nil {
		fmt.Println(err)
		specificError = affected
	}
	return item, specificError
}
