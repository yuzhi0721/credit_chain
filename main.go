package main

import (
	"CreditChain/controller"
	"CreditChain/tools"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./web/index.html", "./web/upload.html")
	r.StaticFS("/static", http.Dir("./web/static"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	// user操作routerGroup
	userGroup := r.Group("/userinfo")
	userGroup.Use(tools.Cors())
	{
		userGroup.GET("/", controller.ShowUser)
		userGroup.POST("/search", controller.SelectUser)
		userGroup.GET("/add", controller.AddUser)
		userGroup.POST("/add", controller.AddUser)
		userGroup.POST("/update")
	}
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(200, "upload.html", nil)
	})
	r.POST("/upload", controller.UploadFile)
	r.Run(":9091").Error()
}
