package models

type BlogEntry struct {
	UserID     int64  `column:"User_ID"`
	AuthorName string `column:"Author_Name"`
	BlogEntry  string `column:"Blog_Text"`
}
