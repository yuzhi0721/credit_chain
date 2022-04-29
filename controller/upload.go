package controller

import (
	"CreditChain/dao"
	"CreditChain/model"
	"CreditChain/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func UploadFile(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(100)
	f, err := c.FormFile("f1")
	if err != nil {
		c.JSON(200, gin.H{
			"code": 4001,
			"err":  "文件获取失败",
		})
		panic(err)
	} else {
		md5Str := tools.MD5(f.Filename)
		path := fmt.Sprintf("D:\\CreditChain\\store\\%s\\", time.Now().Format("20160102"))
		tools.PathExists(path)
		fileName := fmt.Sprintf("%s%d%s", md5Str, randNum, f.Filename)
		path = path + fileName
		id := c.Query("id")
		dao.Connect()
		dao.DB.AutoMigrate(&model.ListStaff{})
		dao.DB.Model(&model.ListStaff{}).Where("id=?", id).Update("imgs", path)
		err = c.SaveUploadedFile(f, path)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 4002,
				"err":  "文件存储失败",
			})
			panic(err)
		} else {
			c.JSON(200, gin.H{
				"message": "上传成功",
			})
		}
	}
}
