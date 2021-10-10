package router

import (
	"go-chat/handler"
	"go-chat/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) *gin.Engine {

	g.Use(middleware.Cors("http://localhost:1234"))
	g.StaticFS("/public", http.Dir("../asset"))
	g.GET("/joke", handler.TakeJoke)

	return g
}
