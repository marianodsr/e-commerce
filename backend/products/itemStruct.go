package products

//Item struct
type Item struct {
	ID       uint `json:"ID"`
	Quantity int  `json:"quantity"`
}

//ItemList struct
type ItemList struct {
	Items []Item `json:"items"`
}
