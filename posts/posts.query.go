package posts

import (
	"websocket/common"
)

var db = common.CreateDBInstance(); 

func GetAllPostsQuery() (*Posts, error){
	session := db.Copy();

	collection := session.DB("Node").C("blogs");

	posts := new(Posts)

	collection.Find(nil).All(posts);

	return posts, nil
}