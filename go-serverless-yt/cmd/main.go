package main

import(
"github.com/M0HTeP/go/go-serverless-yt/pkg/handlers"
"os"
"github.com/aws/aws-lambda-go/events"
"github.com/aws/aws-lambda-go/lambda"
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/dynamodb"
"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var(
	dynaClient dynamodbiface.DynamoDBAPI
)

func main(){
	region := os.Getenv("AWS_REGION")	//	получаем регион
	awsSession, err := session.NewSession(&aws.Config{	//	создаем новую сессию
		Region: aws.String(region)},)	//	записываем регион в переменную

	if err!=nil{
		return
	}
	dynaClient = dynamodb.New(awsSession)	//	начинает новую сессию в dynamodb
	lambda.Start(handler)	//	запуск лямбда-функции
}


const tableName = "go-serverless-yt"	//	название таблицы

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error){	//	хендлер запроса события(?)
		switch req.HTTPMethod{	//	свич(цикл)
		case "GET":
			return handlers.GetUser(req, tableName, dynaClient)	//	получение пользователя
		case "POST":
			return handlers.CreateUser(req, tableName, dynaClient)	//	создание пользователя
		case "PUT":
			return handlers.UpdateUser(req, tableName, dynaClient)	//	изменение данных пользователя
		case "DELETE":
			return handlers.DeleteUser(req, tableName, dynaClient)	//	удаление пользователя
		default:
			return handlers.UnhandledMethod()
		}
}