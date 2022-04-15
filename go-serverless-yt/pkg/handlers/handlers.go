package handlers

import(
	"github.com/M0HTeP/go/go-serverless-yt/pkg/user"
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct{
	ErrorMsg *string `json:"error,omitempty"`
}
	//	функция вывода информации о пользователе
func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(
	*events.APIGatewayProxyResponse, error,
){
	//	получение информации о электронной почте
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)	//	парсим информацию
		if err!= nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})	//	в случае ошибки отправляем статус и описание ошибки
		}
		return apiResponse(http.StatusOK, result)	//	если все нормально - отправляем результат
	}
	//	получаем информацию о ВСЕХ пользователях
	result, err := user.FetchUsers(tableName, dynaClient)
	if err!= nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)

}
	//	функция создания нового пользователя
func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(
	*events.APIGatewayProxyResponse, error,
){
	result, err := user.CreateUser(req, tableName, dynaClient)	//	создаем нового пользователя с заданными параметрами
	if err!=nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{	//	отлавливаем ошибки
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusCreated, result)
}
	//	функция обновления информации о пользователе
func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(
	*events.APIGatewayProxyResponse, error,
){
	result, err := user.UpdateUser(req, tableName, dynaClient)	//	обновляем информацию
	if err!= nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)
}
	//	функция удаления пользователя
func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(
	*events.APIGatewayProxyResponse, error,
){
	err := user.DeleteUser(req, tableName, dynaClient)

	if err!= nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, nil)
}
	//	функция для определения допустимости метода
func UnhandledMethod()(*events.APIGatewayProxyResponse, error){
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}