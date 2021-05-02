package posts

type Post struct {
	Title string `json:"name" bson:"title"`
	Post string `json:"post" bson:"post"`
	Description string `json:"description" bson:"description"`
}

type Posts []Post