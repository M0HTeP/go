package user

import (
	"encoding/json"
	"errors"
	"github.com/M0HTeP/go/go-serverless-yt/pkg/validators"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ( //	описываем возможные ошибки
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorFailedToFetchRecord     = "failed to fetch record"
	ErrorInvalidUserData         = "invalid user data"
	ErrorInvalidEmail            = "invalid email"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item"
	ErrorUserAlreadyExists       = "user.User already exists"
	ErrorUserDoesNotExist        = "user.User does not exist"
)

//	создаем структуру хранения пользователей
type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//	 плучаем инфу о пользователе по электронной почте(вместо ID)
func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {

	input := &dynamodb.GetItemInput{ //	водим почту
		Key: map[string]*dynamodb.AttributeValue{ //	map - преобразует наш ввод
			"email": {
				S: aws.String(email), //	в графу "почта"
			},
		},
		TableName: aws.String(tableName), //	в таблице
	}

	result, err := dynaClient.GetItem(input) //	пытаемся получить пользователя
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord) //	если не удалось - ошибка
	}

	item := new(User)                                       //	запихиваем пользователя в переменную
	err = dynamodbattribute.UnmarshalMap(result.Item, item) //	преобразуем полученный результат
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord) //	если преобразовать не получилось - ошибка
	}
	return item, nil //	возвращаем результат
}

//	получение информации о ВСЕХ пользователях
func FetchUsers(tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*[]User, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName), //	сканируем таблицу
	}

	result, err := dynaClient.Scan(input) //	пытаемся получить результат
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord) //	если не удалось - ошибка
	}
	item := new([]User)                                             //	запихиваем пользователя в переменную
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item) //	преобразуем результат
	return item, nil                                                //	возвращаем результат
}

//	создание нового пользователя
func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	var u User //	создаем переменную со структурой Пользователь

	if err := json.Unmarshal([]byte(req.Body), &u); err != nil { //	перобразуем результат
		return nil, errors.New(ErrorInvalidUserData)
	}
	if !validators.IsEmailValid(u.Email) { //	проверяем почту на валидность
		return nil, errors.New(ErrorInvalidEmail)
	}

	currentUser, _ := FetchUser(u.Email, tableName, dynaClient) //	парсим данные для создания нового пользователя
	if currentUser != nil && len(currentUser.Email) != 0 {      //	проверяем на ошибки
		return nil, errors.New(ErrorUserAlreadyExists)
	}

	av, err := dynamodbattribute.MarshalMap(u) //	преобразуем пользователя в аттрибуты dynamoDB

	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem) //	проверка на ошибки
	}

	input := &dynamodb.PutItemInput{ //	закидываем пользователя в таблицу
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input) //	проверяем заброс на ошибки
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	return &u, nil
}

//	изменение данных пользователя
func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	var u User                                                   //	присваеваем переменной структуру Пользователя
	if err := json.Unmarshal([]byte(req.Body), &u); err != nil { //	проверяем на ошибки наш запрос
		return nil, errors.New(ErrorInvalidEmail)
	}

	currentUser, _ := FetchUser(u.Email, tableName, dynaClient)	//	получаем инфу о пользователе и закидываем ее в переменную
	if currentUser != nil && len(currentUser.Email) == 0 {
		return nil, errors.New(ErrorUserDoesNotExist)	//	отлавливаем ошибки
	}

	av, err := dynamodbattribute.MarshalMap(u)	//	преобразуем пользователя в атрибуты dynamoDB
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{	//	вписываем данные пользователя в таблицу
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)	//	отлавливаем ошибки
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	return &u, nil
}

//	удаление пользователя
func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) error {

	email := req.QueryStringParameters["email"]	//	преобразуем почту
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{	//	ищем пользователя по почте
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),	//	в таблице
	}
	_, err := dynaClient.DeleteItem(input)	//	удаляем пользователя
	if err != nil {
		return errors.New(ErrorCouldNotDeleteItem)
	}

	return nil
}
