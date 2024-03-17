package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	r.GET("home", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "home",
		})
	})
	r.Run()
}