package models

type Book struct {
	ID     string `json:"id" bson:"_id"`
	Title  string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	Year   string `json:"year" bson:"year"`
}
