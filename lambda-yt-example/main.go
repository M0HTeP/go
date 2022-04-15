package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)
    //    создаем структуру данных
type MyEvent struct {
	Name string `json:"what is your name?"`    //    Имя
	Age  int    `json:"How old are you?"`    //    Возраст
}

type MyResponse struct {
	Message string `json:"Answer:"`    //    Вид ответа
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old", event.Name, event.Age)}, nil    //    Отвечаем на сообщение
}

func main() {
	lambda.Start(HandleLambdaEvent)    //    запускаем функцию
}

// aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json
// Не понял, как заставить этот код работать...
