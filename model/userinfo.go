package model

type ListStaff struct {
	Id int32				`json:"id" form:"id"`
	Name string  			`json:"name" form:"name"`						//'姓名'
	Department string  		`json:"department" form:"department"`			//'部门'
	Duties string 			`json:"duties" form:"duties"`					//'职务'
	Tel string  			`json:"tel" form:"tel"`							//'电话号码'
	Num int32  				`json:"num" form:"num"`							//'员工工号'
	Wechat string  			`json:"wechat" form:"wechat"`					//'微信'
	Email string  			`json:"email" form:"email"`						//'邮箱'
	Gender int32  			`json:"gender" form:"gender"`					//'性别，1男2女'
	Native string  			`json:"native" form:"native"`					//'籍贯'
	Education int32  		`json:"education" form:"education"`				//'学历'
	Major string  			`json:"major" form:"major"`						//'专业'
	Id_num string 			`json:"id_num" form:"id_num"`					//'身份证'
	Status int32 			`json:"status" form:"status" gorm:"default:1"`					//'状态1在职2离职'
	Entry_time int64		`json:"entry_time" form:"entry_time"` 			//'入职时间'
	Quit_time  int64  		`json:"quit_time" form:"quit_time" `			//'离职时间'
	Birthday int64 			`json:"birthday" form:"birthday"`
	Graduate_time	int64	`json:"graduate_time" form:"graduate_time"` 	// '毕业时间'
	Input_time int64  		`json:"input_time" form:"input_time" `			//'新增时间'
	Update_time int64   	`json:"update_time" form:"update_time"`			//'最后一次修改时间'
	Userid int32  			`json:"userid" form:"userid"`					//'最后操作人ID'
}

