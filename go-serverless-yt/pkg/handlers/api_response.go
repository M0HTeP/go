package handlers

import(
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)
	//	функция отрисовки действий с таблицей
func apiResponse(status int, body interface{})(*events.APIGatewayProxyResponse, error){
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type":"application/json"}}	//	сам ответ
	resp.StatusCode = status 	//	статус ответа

	stringBody, _ := json.Marshal(body)	//	преобразователь тела запроса
	resp.Body = string(stringBody)	//	конвертор тела запроса
	return &resp, nil
}