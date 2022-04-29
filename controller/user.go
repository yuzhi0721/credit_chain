package controller

import (
	"CreditChain/dao"
	"CreditChain/model"
	"CreditChain/tools"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ShowUser(c *gin.Context) {
	var userslice []model.ListStaff
	err := dao.Connect()
	if err != nil {
		panic(err)
	}
	dao.DB.AutoMigrate(&model.ListStaff{})
	dao.DB.Find(&userslice)
	for _, v := range userslice {
		entryTime := tools.TimeStampToDate(v.Entry_time)
		quitTime := tools.TimeStampToDate(v.Quit_time)
		c.JSON(200, gin.H{
			"id":         v.Id,
			"name":       v.Name,
			"department": v.Department,
			"duties":     v.Duties,
			"tel":        v.Tel,
			"entryTime":  entryTime,
			"mark":       5,
			"quitTime":   quitTime,
			"status":     v.Status,
		})
	}
	defer dao.DB.Close()
}

func SelectUser(c *gin.Context) {
	var user model.ListStaff
	err := dao.Connect()
	if err != nil {
		panic(err)
	}
	key := c.Query("id")
	dao.DB.AutoMigrate(&model.ListStaff{})
	dao.DB.Where("id=?", key).Find(&user)
	if user.Id != 0 {
		c.JSON(200, gin.H{
			"code":          200,
			"name":          user.Name,
			"department":    user.Department,
			"duties":        user.Duties,
			"tel":           user.Tel,
			"num":           user.Num,
			"wechat":        user.Wechat,
			"email":         user.Email,
			"gender":        user.Gender,
			"native":        user.Native,
			"education":     user.Education,
			"major":         user.Major,
			"id_num":        user.Id_num,
			"entry_time":    user.Entry_time,
			"birthday":      user.Birthday,
			"graduate_time": user.Graduate_time,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    4001,
			"message": "用户不存在",
		})
	}
	defer dao.DB.Close()
}

func AddUser(c *gin.Context) {
	var user1 model.ListStaff
	var user2 model.ListStaff
	// 连接数据库，获取数据库连接
	err := dao.Connect()
	if err != nil {
		panic(err)
	}
	dao.DB.AutoMigrate(&model.ListStaff{})
	// POST方法获取到的内容对应解析到user结构体
	c.ShouldBind(&user1)
	defer dao.DB.Close()
	// 通过解析绑定的结构体user1的id_num字段在数据库中查找结果并赋值在结构体user2上
	dao.DB.Where("id_num=?", user1.Id_num).Find(&user2)
	// 对查找赋值的user2进行判断，如果查找结果id_num为""
	fmt.Println(user2)
	if user2.Id != 0 {
		// 返回结果4001，添加失败
		c.JSON(200, gin.H{
			"code":    4001,
			"message": "用户已存在，添加失败",
		})
		//return
	} else {
		// 将前端获取的user1结构体数据内容插入到数据库中
		dao.DB.Create(&user1)
		// 返回前端结果code200，json
		c.JSON(200, gin.H{
			"code":    200,
			"message": "添加成功",
		})
	}

}

// 更新用户信息，用户名,身份证号码不能更改

func ShowUpdate(c *gin.Context) {
	var user model.ListStaff
	err := dao.Connect()
	if err != nil {
		panic(err)
	}
	id := c.Query("id")
	dao.DB.AutoMigrate(&model.ListStaff{})
	dao.DB.Where("id=?", id).Find(&user)
	if user.Id != 0 {
		c.JSON(200, gin.H{
			"code":          200,
			"name":          user.Name,
			"department":    user.Department,
			"duties":        user.Duties,
			"tel":           user.Tel,
			"num":           user.Num,
			"wechat":        user.Wechat,
			"email":         user.Email,
			"gender":        user.Gender,
			"native":        user.Native,
			"education":     user.Education,
			"major":         user.Major,
			"id_num":        user.Id_num,
			"entry_time":    user.Entry_time,
			"birthday":      user.Birthday,
			"graduate_time": user.Graduate_time,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    4001,
			"message": "用户不存在",
		})
	}
	defer dao.DB.Close()
}

func Update(c *gin.Context) {
	//获取前端hidden隐藏传参用户ID
	userId := c.Query("id")
	err := dao.Connect()
	if err != nil {
		panic(err)
	}
	dao.DB.AutoMigrate(&model.ListStaff{})
	defer dao.DB.Close()
	var user model.ListStaff
	var userReceive model.ListStaff
	// 获取前端传过来的数据内容绑定在结构体user上
	c.ShouldBind(&user)
	//查找身份证等于传值且id不等于传值的数据赋值给userReceive
	dao.DB.Where("id_num=? AND id<>?", user.Id_num, userId).Find(&userReceive)
	fmt.Println(userReceive)
	if userReceive.Name == "" {
		fmt.Println(user)
		//err := dao.DB.Model(&user).Updates(model.ListStaff{Name:user.Name,Department: user.Department,Duties: user.Duties,Tel: user.Tel,Num: user.Num,Wechat: user.Wechat,Email: user.Email,Gender: user.Gender,Native: user.Native,Education: user.Education,Major: user.Major,Id_num: user.Id_num,Entry_time: user.Entry_time,Birthday: user.Birthday,University: user.University,Graduate_time: user.Graduate_time}).Error
		err := dao.DB.Save(&user).Error
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

func DelUser() {
}
