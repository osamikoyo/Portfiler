package models

type Review struct {
	Username string `bson:"username"`
	Content  string `bson:"review_content"`
}

type Portfolio struct {
	ID      uint     `bson:"id"`
	Title   string   `bson:"title"`
	Content string   `bson:"content"`
	Reviews []Review `bson:"reviews"`
}
