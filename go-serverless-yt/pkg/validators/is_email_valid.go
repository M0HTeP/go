package validators

import "regexp"

func IsEmailValid(email string) bool {	//	проверка электронной почты на валидность
		//	прописываем какие символы должна содержать почта
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]{1,64}@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) < 3 || len(email) > 254 || !rxEmail.MatchString(email){	//	ограничение по длине
		return false
	}

	return true
}