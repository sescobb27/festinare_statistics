package models

type Discount struct {
	Categories []*Category `json:"categories"`
	Hashtags   []*string   `json:"hashtags"`
}
