package posts

func GetAllPostService() (*Posts, error) {
	posts, err := GetAllPostsQuery();
	return posts, err;
}