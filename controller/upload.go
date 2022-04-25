package controller

import (
	"CreditChain/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func UploadFile(c *gin.Context){
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(100)
	f,err := c.FormFile("f1")
	if err != nil{
		c.JSON(400,gin.H{
			"err":err.Error(),
		})
	}else{
		md5Str := tools.MD5(f.Filename)
		path := fmt.Sprintf("D:\\CreditChain\\store\\%s\\",time.Now().Format("20160102"))
		tools.PathExists(path)
		fileName := fmt.Sprintf("%s%d%s",md5Str,randNum,f.Filename)
		path = path + fileName
		err = c.SaveUploadedFile(f,path)
		if err !=nil{
			c.JSON(200,gin.H{
				"err":err.Error(),
			})
		}else {
			c.JSON(200,gin.H{
				"message":"OK",
			})
		}
	}
}
