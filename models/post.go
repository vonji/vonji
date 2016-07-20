package models

type Post struct {
	User User
	UserID uint
	Content string
	Comments []Comment
}
