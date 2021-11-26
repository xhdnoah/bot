package router

import (
	"go-chat/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) *gin.Engine {

	g.StaticFS("/public", http.Dir("../asset"))
	g.GET("/joke", handler.TakeJoke)

	return g
}
