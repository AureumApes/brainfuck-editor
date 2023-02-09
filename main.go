package main

import (
	router2 "github.com/AureumApes/brainfuck-editor/router"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", router2.Index)
	router.POST("/build", router2.Build)
	router.POST("/save", router2.Save)
	router.Run(":8080")
}
