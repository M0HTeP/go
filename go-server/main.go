package main //

import ( // импортируем необходимые модули
	"fmt"      // для вывода сообщений
	"log"      // для логгирования ошибок
	"net/http" // для работы с html файлами
)

// создаем функцию для работы с формой (заполнение имени и адреса))
func formHandler(w http.ResponseWriter, r *http.Request) {
	// в случае возникновения ошибки - выводим, что с модулем ParseForm произошла ошибка и ее номер
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}
	// если ошибок нет - выводим сообщение об этом и затем выводим Имя и Адрес
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// создаем функцию "приветствия"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// если ссылка введена не верно, то выводим ошибку "страница не найдена"
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// если используется метод отличный от "получения", то выводится ошибка о неподдерживаемом методе
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	// если все нормально - выводится приветствие
	fmt.Fprintf(w, "hello")
}

func main() { // основная функция
	fileserver := http.FileServer(http.Dir("./static")) // задаем директорию сервера
	http.Handle("/", fileserver)                        //
	http.HandleFunc("/form", formHandler)               // функция form, которая описывается в графе /form в файле form.html
	http.HandleFunc("/hello", helloHandler)             // функция hello, которая описывается в графе /hello в файле form.html

	fmt.Println("Starting server at port 8080\n") // выводим сообщение о запуске сервера
	// в случае возникновения ошибки - записываем ее в лог
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
