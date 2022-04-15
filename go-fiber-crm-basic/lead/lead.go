package lead

import (
	"github.com/M0HTeP/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//    структура нашей БД
type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

//    функция получения всех записей БД
func GetLeads(c *fiber.Ctx) {
	db := database.DBConn //    соединяемся с БД
	var leads []Lead      //    переменная для хранения данных БД
	db.Find(&leads)       //    ищем значения
	c.JSON(leads)         //    выводим значения
}

//    функция получения информации об определенном ID
func GetLead(c *fiber.Ctx) {
	id := c.Params("id") //    выцепляем ID из запроса
	db := database.DBConn
	var lead Lead	//    переменная для хранения объекта БД
	db.Find(&lead, id) //    ищем объект БД с нужным ID
	c.JSON(lead)
}

//    функция создания нового объекта БД
func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)                          //    переменная с новым объектом
	if err := c.BodyParser(lead); err != nil { //    парсим тело запроса
		c.Status(503).Send(err) //    в случае ошибки отправляем статус "ошибка"
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

//    функция удаления объекта БД
func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)	//    Ищем первое совпадение объекта БД с нужным ID
	if lead.Name == "" {
		c.Status(500).Send("No lead found with ID")	//    если такого ID не существует так и говорим
		return
	}
	db.Delete(&lead)	//    удаляем найденный объект БД
	c.Send("Lead successfully deleted")	//    отчитываемся о выполнении

}
