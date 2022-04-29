package controller

import (
	"CreditChain/dao"
	"CreditChain/model"
	"github.com/gin-gonic/gin"
)

func CompanyInfo(c *gin.Context) {
	key := c.Query("id")
	err := dao.Connect()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	var company model.CompanyEssinfos
	dao.DB.AutoMigrate(&model.CompanyEssinfos{})
	dao.DB.Where("id=?", key).Find(&company)
	c.JSON(200, gin.H{
		"code":        200,
		"companyName": company.Company,
		"wallet_addr": company.Wallet_addr,
		"uscc":        company.Uscc,
	})

}

func CompanyInfoUpdate(c *gin.Context) {
	//获取前端hidden隐藏传参公司主键ID
	key := c.Query("id")
	err := dao.Connect()
	if err != nil {
		panic(err)
	}
	dao.DB.AutoMigrate(&model.CompanyEssinfos{})
	defer dao.DB.Close()
	var company model.CompanyEssinfos
	var companyReceive model.CompanyEssinfos
	// 获取前端传过来的数据内容绑定在结构体user上
	c.ShouldBind(&company)
	//查找身份证等于传值且id不等于传值的数据赋值给userReceive
	dao.DB.Where("uscc=? AND id<>?", company.Uscc, key).Find(&companyReceive)
	if companyReceive.Company == "" {
		//err := dao.DB.Model(&user).Updates(model.ListStaff{Name:user.Name,Department: user.Department,Duties: user.Duties,Tel: user.Tel,Num: user.Num,Wechat: user.Wechat,Email: user.Email,Gender: user.Gender,Native: user.Native,Education: user.Education,Major: user.Major,Id_num: user.Id_num,Entry_time: user.Entry_time,Birthday: user.Birthday,University: user.University,Graduate_time: user.Graduate_time}).Error
		err := dao.DB.Save(&company).Error
		if err != nil {
			c.JSON(200, gin.H{
				"code":    4002,
				"message": "数据库意外错误修改失败",
			})
			panic(err)
		} else {
			c.JSON(200, gin.H{
				"code":    200,
				"message": "修改成功",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"code":    4001,
			"message": "修改失败，信息重复",
		})
	}
}
