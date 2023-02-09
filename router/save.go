package router

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Save(ctx *gin.Context) {
	name, _ := ctx.GetQuery("name")
	body := ctx.Request.Body
	bodyBytes, _ := io.ReadAll(body)

	os.WriteFile("./dist/"+name+".bf", bodyBytes, os.ModePerm)
}
