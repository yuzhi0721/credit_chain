package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var (
	DB *gorm.DB
)
func Connect ()(err error){
	DB ,err = gorm.Open("mysql","root:123123@(42.192.149.62:3306)/CreditChain")
	if err !=nil{
		panic(err)
	}
	return DB.DB().Ping()
}

