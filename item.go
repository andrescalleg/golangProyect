package main

type Item struct {
	id     int32
	name   string
	value  float64
}

type CreateItem struct {
    ID      int `json:"id"`
    Name    string `json:"name"`
    Value   float64 `json:"value"`
    CartID  int `json:"cart_id"`
    Amount  int `json:"amount"`
}
