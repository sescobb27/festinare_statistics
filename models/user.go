package models

type User struct {
	Locations  *Location   `json:"locations"`
	Discounts  []*Discount `json:"discounts"`
	Categories []*Category `json:"categories"`
	Mobile     *Mobile     `json:"mobile"`
}
