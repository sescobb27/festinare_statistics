package models

type Client struct {
	Categories []*Category `json:"categories"`
	Locations  []*Location `json:"locations"`
	Discounts  []*Discount `json:"discounts"`
	Addresses  []*string   `json:"addresses"`
	Name       string      `json:"name"`
}
