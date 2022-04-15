package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//    создаем структуру
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

//    создаем переменную для хранения данных
var mg MongoInstance

//	создаем константы
const dbName = "fiber-hrms"                            //	имя БД
const mongoURI = "mongodb://localhost:27017/" + dbName //	ссылка на адрес БД + имя БД

//	создаем структуру "Работники"
type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"` //	omitempty - необязательное поле
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

//	функция соединения и перехвата ошибок
func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))      //	создаем нового клиента, проходящего по ссылке на нашу БД и переменную ошибки
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) //	для обрыва соединения с БД после таймаута, если ничего не происходит
	defer cancel()

	err = client.Connect(ctx)     //	заносим в переменную для ошибки контекстную информацию о соединении
	db := client.Database(dbName) //	заносим в переменную клиента, подсоединившегося к БД

	if err != nil {
		return err //	в случае ошибки - возвращаем ее
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err) //	логгируем ошибку, если она есть
	}
	app := fiber.New() //	создаем новый инстанс

	app.Get("/employee", func(c *fiber.Ctx) error { //	ищем список работников в БД по контексту сообщения

		query := bson.D{{}} //	запрос

		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query) //	ищем в графе "работник" информацию из контекста запроса
		if err != nil {
			return c.Status(500).SendString(err.Error()) //	в случае ошибки возвращаем ее
		}

		var employees []Employee = make([]Employee, 0) //	создаем список работников

		if err := cursor.All(c.Context(), &employees); err != nil { //
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(employees) //	возвращаем в формате JSON
	})

	app.Post("/employee", func(c *fiber.Ctx) error {
		collection := mg.Db.Collection("employees")

		employee := new(Employee) //	создаем нового сотрудника

		if err := c.BodyParser(employee); err != nil { //	парсим данные нового сотрудника
			return c.Status(400).SendString(err.Error())
		}

		employee.ID = "" //	пустой ID

		insertionResult, err := collection.InsertOne(c.Context(), employee) //	вставляем документ, содержащий данные о работнике в коллекцию

		if err != nil {
			return c.Status(500).SendString(err.Error()) //	если ошибка - кричим об этом
		}

		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}} //	проверка на наличие ошибок с т.зрения MongoDB
		createdRecord := collection.FindOne(c.Context(), filter)          //	создаем запись

		createdEmployee := &Employee{}        //	переменная содержащая нового работника(ссылку на него)
		createdRecord.Decode(createdEmployee) //	декодим запись о новом работнике

		return c.Status(201).JSON(createdEmployee) //	возвращаем статус создания новой записи

	})
	//	функция изменения данных о работнике
	app.Put("/employee/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id") //	выцепляем ID из запроса

		employeeID, err := primitive.ObjectIDFromHex(idParam) //	(?)

		if err != nil {
			return c.SendStatus(400) //	в случае ошибки - кричим об этом
		}

		employee := new(Employee) //	вся инфа о работнике, которую нужно вставить в замен существующей

		if err := c.BodyParser(employee); err != nil { //	парсим инфу(тело запроса) о работнике
			return c.Status(400).SendString(err.Error())
		}

		query := bson.D{{Key: "_id", Value: employeeID}} //	создаем запрос в БД

		update := bson.D{ //	команда Обновления записи ня языке SQL
			{Key: "$set",
				Value: bson.D{
					{Key: "name", Value: employee.Name},
					{Key: "age", Value: employee.Age},
					{Key: "salary", Value: employee.Salary},
				},
			},
		}

		err = mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err() //	переменная содержащая ошибку

		if err != nil {
			if err == mongo.ErrNoDocuments { //	если нет информации об ошибке в документации MongoDB
				return c.SendStatus(400)
			}
			return c.SendStatus(500)
		}

		employee.ID = idParam //	подтягиваем ID из запроса

		return c.Status(200).JSON(employee)

	})
	//	функция удаления сотрудника
	app.Delete("/employee/:id", func(c *fiber.Ctx) error {

		employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))	//	выцепляем ID из запроса

		if err != nil {
			return c.SendStatus(400)	//	в случае ошибки отправляем статус
		}

		query := bson.D{{Key: "_id", Value: employeeID}}	//	создаем запрос
		result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), &query)	//	удаляем запись из БД

		if err != nil {
			return c.SendStatus(500)	//	если ошибка
		}

		if result.DeletedCount < 1 {
			return c.SendStatus(404)	//	если удалили менее 1 записи(404 - не нашли запись)
		}

		return c.Status(200).JSON("record deleted")	//	сообщаем об успехе

	})

	log.Fatal(app.Listen(":3000"))	//	логируем ошибки на нашем порту
}
