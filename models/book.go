package models

type Book struct {
	BookID int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}
