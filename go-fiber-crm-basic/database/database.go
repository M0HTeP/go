package database


import(
"github.com/jinzhu/gorm"
_"github.com/jinzhu/gorm/dialects/sqlite"

)


var(
	DBConn *gorm.DB 	//    переменная, с помощью которой соединяемся с БД
)