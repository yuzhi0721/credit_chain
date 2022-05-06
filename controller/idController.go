package controller

import (
	"CreditChain/dao"
	"CreditChain/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv18/pkg/result"
	"github.com/liuhongdi/digv18/service"
	//"github.com/liuhongdi/digv18/service"
)

type IdController struct{}

func NewIdController() IdController {
	return IdController{}
}

//存储验证码的结构
type CaptchaResult struct {
	Id         string `json:"id"`
	Base64Blob string `json:"base_64_blob"`
	//VerifyValue string `json:"code"`
}

// 生成图形化验证码
func (a *IdController) GetOne(ctx *gin.Context) {
	resultRes := result.NewResult(ctx)
	id, b64s, err := service.CaptMake()
	if err != nil {
		resultRes.Error(1, err.Error())
	}
	captchaResult := CaptchaResult{
		Id:         id,
		Base64Blob: b64s,
	}
	resultRes.Success(captchaResult)
	return
}

//验证captcha是否正确
func (a *IdController) Verify(c *gin.Context) {
	id := c.PostForm("id")
	capt := c.PostForm("capt")
	username := c.PostForm("username")
	password := c.PostForm("password")
	resultRes := result.NewResult(c)
	if id == "" || capt == "" {
		resultRes.Error(400, "param error")
	}
	if service.CaptVerify(id, capt) == true {
		dao.Connect()
		defer dao.DB.Close()
		dao.DB.AutoMigrate(&model.User{})
		var user model.User
		dao.DB.Where("username=?", username).Find(&user)
		fmt.Println(user.Password)
		if user.Password == password {
			resultRes.Success("登录成功")
			fmt.Println("login successful")
		} else {
			fmt.Println("login failed")
			resultRes.Error(2, "密码错误")
		}
	} else {
		resultRes.Error(1, "验证码错误")
	}
}
