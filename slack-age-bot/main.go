package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
)

// создаем функцию, которая будет выводить сообщение бота
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) { // передаем событие в канал
	for event := range analyticsChannel { //  проходимся по событиям в цикле
		fmt.Println("Command Events") // выводим сообщение
		fmt.Println(event.Timestamp)  // проставляем время
		fmt.Println(event.Command)    // выводим команду
		fmt.Println(event.Parameters) // выводим параметры команды
		fmt.Println(event.Event)      // выводим событие
		fmt.Println()                 // пустая строка(?)
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3398622637444-3420057801264-RTQANQrkmWS1WLgo3AcLlDei")                                         //  токен бота
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03B7L3NV55-3396447905939-da05fb9c6ef180444ad6f923ffcefbd877bfb434c63633606a54f109837e5598") //  токен приложения

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN")) // создание бота

	go printCommandEvents(bot.CommandEvents()) //  (?)

	bot.Command("my yob is <year>", &slacker.CommandDefinition{ // задаем боту команду
		Description: "yob calculator", //  описание команды
		Example:     "my yob is 2020", // пример
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) { // создаем хендлер для контекста передаваемой команды
			year := request.Param("year")  // переменная "год" берется из сообщения боту
			yob, err := strconv.Atoi(year) // переводится в строку
			if err != nil {
				println("error") // в случае ошибки выводится
			}
			age := 2022 - yob                  //  математические вычисления =)
			r := fmt.Sprintf("age is %d", age) //  форматируем вывод
			response.Reply(r)                  // отвечаем
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err) // ошибку записываем в лог
	}
}
