package model

type Recipe struct {
	Id          uint32   `json:"id"`
	Name        string   `json:"name"`
	Ingridients []string `json:"ingridients"`
	Description string   `json:"description"`
}
