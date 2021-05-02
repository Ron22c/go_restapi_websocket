package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"

	"websocket/posts"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WELCOME TO THE CHATTER");
}

func Contains(slice []string, val string) bool{
	for _, value := range slice {
		if value == val {
			return true
		}
	}

	return false;
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {

	// COmmenting the code because middlewire recoverer does the same thing
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(r.Header["Connection"]);
	// 		fmt.Fprintf(w, "error this is not a rest endpoint")
	// 	}
	// }()

	upgrader.CheckOrigin = func(r *http.Request) bool {return true};

	ws, err := upgrader.Upgrade(w, r, nil);

	if err != nil {
		fmt.Println(err);
	}

	HandleConnection(ws);
}

func HandleConnection(con *websocket.Conn) {
	for {
		msgType, msg, err := con.ReadMessage();

		if err != nil {
			fmt.Println(err);
		}

		fmt.Println(string(msg), msgType);

		if err := con.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err);
		}
	}
}

func CreateRoutes() http.Handler{
	r := chi.NewRouter();
	r.Use(middleware.Logger);
	r.Use(middleware.Recoverer)

	r.Get("/", GetHomePage);
	r.Get("/ws", WebSocketHandler);

	r.Mount("/post", posts.NewPostRouter());

	return r;
}

func main() {
	handler := CreateRoutes()
	fmt.Println(http.ListenAndServe(":8080", handler));
}