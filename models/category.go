package models

var CATEGORIES = [3]string{"Bar", "Disco", "Restaurant"}

type Category struct {
	Name string `json:"name"`
}
