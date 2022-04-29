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
		//  根据服务器存储路径进行修改
		path := fmt.Sprintf("D:\\CreditChain\\Userimages\\%s\\", time.Now().Format("20160102"))
		tools.PathExists(path)
		fileName := fmt.Sprintf("%s%d%s", md5Str, randNum, f.Filename)
		path = path + fileName
		err = c.SaveUploadedFile(f, path)
		id := c.Query("id")
		if err != nil {
			c.JSON(200, gin.H{
				"code": 4002,
				"err":  "文件存储失败",
			})
			panic(err)
		} else {
			dao.Connect()
			defer dao.DB.Close()
			dao.DB.AutoMigrate(&model.ListStaff{})
			dao.DB.Model(&model.ListStaff{}).Where("id=?", id).Update("imgs", path)
			c.JSON(200, gin.H{
				"message": "上传成功",
			})
		}
	}
}

func UploadLOGO(c *gin.Context) {
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
		//  根据服务器存储路径进行修改
		path := fmt.Sprintf("D:\\CreditChain\\LOGO\\%s\\", time.Now().Format("20160102"))
		tools.PathExists(path)
		fileName := fmt.Sprintf("%s%d%s", md5Str, randNum, f.Filename)
		path = path + fileName
		id := c.Query("id")
		err = c.SaveUploadedFile(f, path)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 4002,
				"err":  "文件存储失败",
			})
			panic(err)
		} else {
			dao.Connect()
			defer dao.DB.Close()
			dao.DB.AutoMigrate(&model.CompanyEssinfos{})
			dao.DB.Model(&model.CompanyEssinfos{}).Where("id=?", id).Update("LOGO", path)
			c.JSON(200, gin.H{
				"message": "上传成功",
			})
		}
	}
}

func UploadUSCC(c *gin.Context) {
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
		//  根据服务器存储路径进行修改
		path := fmt.Sprintf("D:\\CreditChain\\USCC\\%s\\", time.Now().Format("20160102"))
		tools.PathExists(path)
		fileName := fmt.Sprintf("%s%d%s", md5Str, randNum, f.Filename)
		path = path + fileName
		id := c.Query("id")
		err = c.SaveUploadedFile(f, path)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 4002,
				"err":  "文件存储失败",
			})
			panic(err)
		} else {
			dao.Connect()
			defer dao.DB.Close()
			dao.DB.AutoMigrate(&model.CompanyEssinfos{})
			dao.DB.Model(&model.CompanyEssinfos{}).Where("id=?", id).Update("uscc", path)
			c.JSON(200, gin.H{
				"message": "上传成功",
			})
		}
	}
}
