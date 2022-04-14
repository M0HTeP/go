package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// создание базы данных
var (
	db *gorm.DB
)

// функция подключения к базе данных
func Connect() {
	d, err := gorm.Open("mysql", "MoHTeP:MOHTEP@/simplerest?charset=utf8&parseTime=True&loc=Local") // данные для подключения. err - ошибка
	if err != nil {
		panic(err) // остановка функции
	}
	db = d // если все нормально - база данных подключена и назначена переменной
}

// функция для получения базы данных из переменной
func GetDB() *gorm.DB {
	return db
}
