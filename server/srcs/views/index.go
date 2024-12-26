package views

import "github.com/gin-gonic/gin"

func Index(context *gin.Context) {
	context.HTML(200, "index.html", nil)
}
