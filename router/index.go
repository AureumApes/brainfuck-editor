package router

import (
	_ "embed"
	"github.com/gin-gonic/gin"
)

//go:embed index.html
var IndexFile string

func Index(ctx *gin.Context) {
	ctx.Writer.WriteString(IndexFile)
}
