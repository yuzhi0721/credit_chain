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
	r.LoadHTMLFiles("./web/index.html", "./web/upload.html", "./web/getcapt.html")
	r.StaticFS("/static", http.Dir("./web/static"))

	r.GET("/login", controller.Login)

	idc := controller.NewIdController()
	r.GET("/id/getone", idc.GetOne)
	r.POST("/id/verify", idc.Verify)

	// user操作routerGroup
	userGroup := r.Group("/userinfo")
	userGroup.Use(tools.Cors())
	{
		userGroup.GET("/", controller.ShowUser)
		userGroup.POST("/show", controller.SelectUser)
		userGroup.GET("/add", controller.AddUser)
		userGroup.POST("/add", controller.AddUser)
		userGroup.GET("/update", controller.ShowUpdate)
		userGroup.POST("/update", controller.Update)
		userGroup.POST("/upload", controller.UploadFile)
	}

	companyGroup := r.Group("/company")
	{
		companyGroup.GET("/baseinfo", controller.CompanyInfo)
		companyGroup.POST("/baseinfo", controller.CompanyInfoUpdate)
		companyGroup.POST("/uploadLOGO", controller.UploadLOGO)
		companyGroup.POST("/uploadUSCC", controller.UploadUSCC)
	}

	r.Run(":9091").Error()
}
