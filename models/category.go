package models

const CATEGORIES = []string{"Bar", "Disco", "Restaurant"}

type Category struct {
	Name string `json:"name"`
}
