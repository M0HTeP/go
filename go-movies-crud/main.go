package main //

import (
	"encoding/json" // для работы с форматом json
	"fmt"           // для вывода сообщений в консоль
	"log"           // для ведения логов ошибок
	"math/rand"     // для рандомной генерации ID
	"net/http"      // для работы с http
	"strconv"       // для конвертации в строку

	"github.com/gorilla/mux" // внешний модуль для работы с http
)

// создаем структуру для обозначения того, какие поля будут использоваться для хранения видиотеки
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Так же создаем структуру для хранения режиссеров
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie // создаем список фильмов

// создаем функцию для получения списка ВСЕХ фильмов. w - ответственен за ответ на запрос и r - указатель на запрос, который мы отправляем на сайт
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //
	json.NewEncoder(w).Encode(movies)                  // кодируем ответ в формат json
}

// создаем функцию удаления фильма по его ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // возвращаем значение переменной, если такая переменная есть
	for index, item := range movies {
		if item.ID == params["id"] { // если в списке существует фильм с таким же ID, как в запросе
			movies = append(movies[:index], movies[index+1:]...) // удаляем фильм с помощью среза, а именно: Пересоздается список исключая нужный ID.
			break
		}
	}
	json.NewEncoder(w).Encode(movies) // перекодируем в формат json
}

// создаем функцию получения данных о фильме с запрашиваемым ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // возвращаем значение переменной, если такая переменная есть
	for _, item := range movies {
		if item.ID == params["id"] { // если в списке существует фильм с таким же ID, как в запросе
			json.NewEncoder(w).Encode((item)) // перекодируем ответ в формат json
			return
		}
	}
}

// создаем функцию, которая будет добавлять в фильмотеку фильм с указанными нами данными
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie                               // создаем переменную структуры нашей фильмотеки
	_ = json.NewDecoder(r.Body).Decode(&movie)    // декодируем тело нашего запроса из json в формат списка фильмов
	movie.ID = strconv.Itoa(rand.Intn(100000000)) // рандомим значение ID для добавляемого фильма (по идее для этого нужно последовательно ID назначать, но на все воля автора) и конвертируем его в строку
	movies = append(movies, movie)                // добавляем новый фильм
	json.NewEncoder(w).Encode(movie)              // кодируем в формат json
}

// создаем функцию, которая будет обновлять информацию о существующем фильме
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // возвращаем значение переменной, если такая переменная есть
	for index, item := range movies {
		if item.ID == params["id"] { // если в списке существует фильм с таким же ID, как в запросе
			movies = append(movies[:index], movies[index+1:]...) // удаляем фильм с помощью среза, а именно: Пересоздается список исключая нужный ID.
			var movie Movie                                      // создаем переменную структуры нашей фильмотеки
			_ = json.NewDecoder(r.Body).Decode(&movie)           // декодируем тело нашего запроса из json в формат списка фильмов
			movie.ID = params["id"]                              // создаем фильм с ID из нашего запроса
			movies = append(movies, movie)                       // добавляем остальную информацию к созданному фильму
			json.NewEncoder(w).Encode(movie)                     // кодируем в формат json
			return
		}
	}
}

func main() {
	r := mux.NewRouter() //

	// Добавляем 2 фильма в нашу фильмотеку
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")           // функция получения списков ВСЕХ фильмов
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")       // функция получения фильма по номеру его ID
	r.HandleFunc("/movies", createMovie).Methods("POST")        // функция добавления новго фильма
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    // функция изменения информации у фильма по его ID
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") // функция удаления фильма из списка по его ID

	fmt.Printf("Starting server at port 8000") // выводим в консоль на каком порту запускаем
	log.Fatal(http.ListenAndServe(":8000", r)) // в случае ошибки на нашем порту - записываем ее в лог
}
