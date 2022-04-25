package controller

import (
	"CreditChain/dao"
	"CreditChain/model"
	"CreditChain/tools"
	"github.com/gin-gonic/gin"
)


func ShowUser(c *gin.Context){
		var userslice []model.ListStaff
		err := dao.Connect()
		if err != nil{
			panic(err)
		}
		dao.DB.AutoMigrate(&model.ListStaff{})
		dao.DB.Find(&userslice)
		for _,v := range userslice{
			 entryTime := tools.TimeStampToDate(v.Entry_time)
			 quitTime := tools.TimeStampToDate(v.Quit_time)
			c.JSON(200,gin.H{
				"id":v.Id,
				"name":v.Name,
				"department":v.Department,
				"duties":v.Duties,
				"tel":v.Tel,
				"entryTime":entryTime,
				"mark":5,
				"quitTime":quitTime,
				"status":v.Status,
			})
		}
		defer dao.DB.Close()
}

func SelectUser(c *gin.Context){
	var user model.ListStaff
	err := dao.Connect()
	if err != nil{
		panic(err)
	}
	key :=  c.Query("id_num")
	dao.DB.AutoMigrate(&model.ListStaff{})
	dao.DB.Where("id_num=?",key).Find(&user)
	if user.Id != 0{
		c.JSON(200,user)
	}else {
		c.JSON(200,gin.H{
			"message":"用户不存在",
		})
	}
	defer dao.DB.Close()
}

func AddUser(c *gin.Context){
	var user1 model.ListStaff
	var user2 model.ListStaff
	// 连接数据库，获取数据库连接
	err := dao.Connect()
	if err != nil{
		panic(err)
	}
	dao.DB.AutoMigrate(&model.ListStaff{})
	// POST方法获取到的内容对应解析到user结构体
	c.ShouldBind(&user1)
	defer dao.DB.Close()
	// 通过解析绑定的结构体user1的name字段在数据库中查找结果并赋值在结构体user2上
	dao.DB.Where("id_num=?",user1.Id_num).Find(&user2)
	// 对查找赋值的user2进行判断，如果查找结果name为""
	if user2.Id_num == "" {
		// 将前端获取的user1结构体数据内容插入到数据库中
		dao.DB.Create(&user1)
		// 返回前端结果code200，json
		c.JSON(200,gin.H{
			"message":"添加成功",
		})
	}else{
		// 返回结果4001，添加失败
		c.JSON(4001,gin.H{
			"message":"添加失败、用户已存在",
		})
	}

}
func DelUser(){
}
