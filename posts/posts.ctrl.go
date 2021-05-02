package posts

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := GetAllPostService();

	if err != nil {
		fmt.Fprintf(w, "ERROR");
	}

	data, err := json.Marshal(posts);

	if err != nil {
		fmt.Fprintf(w, "ERROR");
	}
	
	fmt.Fprintf(w, string(data));
}