package main

type Item struct {
	id     int32
	name   string
	value  float64
}

type CreateItem struct {
	id     int32
	name   string
	value  float64
	cartId int32
	amount int32
}
