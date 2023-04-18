package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Index(ctx *gin.Context) {
	time.Sleep(time.Second)
	ctx.String(http.StatusOK, "Welcome Gin Register")
}
