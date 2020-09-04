package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func getAllCart(db *sql.DB) {
	sqlStatement := fmt.Sprintf(`
		SELECT 
		cart.cart_id, cart.user_name, cart_items.item_amount, item.item_id, item.item_name, item.item_value 
		FROM 
		cart_items, cart, item 
		WHERE 
		cart.cart_id = %d AND cart_items.cart_id = cart.cart_id AND item.item_id = cart_items.item_id;`, 1)
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
		fmt.Println(err)
	}

}

func listDbItems(db *sql.DB) ([]Item, error) {

	var specificError error
	sqlStatement := fmt.Sprintf(`
		select 
		item_id, item_name, item_value  
		from 
		public.item`)

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
	INSERT INTO 
	public.cart(
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

func findCart(cartID int, db *sql.DB) (Cart, error) {
	sqlStatement := fmt.Sprintf(`
		SELECT 
		cart_id, user_name 
		FROM 
		public.cart
		where cart_id = %d ;`, cartID)
	fmt.Println(sqlStatement)

	var cart Cart
	row := db.QueryRow(sqlStatement)
	err := row.Scan(&cart.id, &cart.UserName)
	return cart, err

}

func addItemToCart(newItem CreateItem, db *sql.DB) (Item, error) {
	var item Item
	cart, err := findCart(newItem.CartID, db)
	if err != nil {
		fmt.Println(err)
		return item, err
	}
	if cart.UserName == "" {
		fmt.Println("No cart to add item")
		return item, err
	}
	sqlStatement := fmt.Sprintf(`
	INSERT INTO public.item(
		item_id, item_name, item_value)
		VALUES (%d, '%s', %f);
		`, newItem.ID, newItem.Name, newItem.Value)
	fmt.Println(sqlStatement)

	_, err = db.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err)
		return item, err
	}

	sqlStatementInsert := fmt.Sprintf(`
	INSERT INTO public.cart_items(
		item_id, cart_id, item_amount)
		VALUES (%d, '%d', %d);
		`, newItem.ID, newItem.CartID, newItem.Amount)
	fmt.Println(sqlStatementInsert)

	_, err = db.Exec(sqlStatementInsert)
	if err != nil {
		fmt.Println(err)
		return item, err
	}
	return item, nil
}
