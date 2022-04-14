package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) { // получение запроса
	if body, err := ioutil.ReadAll(r.Body); err == nil { // чтение "тела" запроса
		if err := json.Unmarshal([]byte(body), x); err != nil { // в случае, если в запросе ошибок нет - парсим данные в формате json
			return
		}
	}
}
