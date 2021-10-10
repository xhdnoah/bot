package main

import (
	"go-chat/handler"
	"go-chat/model"
	"go-chat/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Fail to load .env")
	}

	model.DB.Init()
	log.Printf("The database inited.")
	port := os.Getenv("PORT")

	g := gin.New()

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "ai", func(s socketio.Conn, msg string) {
		var ai = make(chan string)
		go handler.HandleAi(ai, msg)
		s.Emit("ai", <-ai)
	})

	server.OnEvent("/", "define", func(s socketio.Conn, msg string) {
		var d = make(chan string, 2)
		go handler.HandleDefine(d, msg)
		for {
			defi, ok := <-d
			if ok {
				s.Emit("define", defi)
			}
		}
	})

	server.OnEvent("/", "weather", func(s socketio.Conn, msg string) {
		var w = make(chan string)
		go handler.HandleWeather(w)
		s.Emit("weather", <-w)
	})

	server.OnEvent("/", "stock", func(s socketio.Conn, msg string) {
		var st = make(chan string)
		go handler.HandleStock(st, msg)
		s.Emit("stock", <-st)
	})

	server.OnEvent("/", "chat", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	g.GET("/socket.io/*any", gin.WrapH(server))
	g.POST("/socket.io/*any", gin.WrapH(server))

	router.Load(g)
	if err := g.Run(`:` + port); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
