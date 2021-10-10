package handler

import (
	"go-chat/model"

	"github.com/gin-gonic/gin"
)

func TakeJoke(c *gin.Context) {
	var joke model.JokeModel
	joke = joke.Take()
	SendResponse(c, nil, joke)
}
