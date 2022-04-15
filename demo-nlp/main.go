package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/krognol/go-wolfram"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"

	witai "github.com/wit-ai/wit-go/v2"
)

//	подключаемся к Вольфраму
var wolframClient *wolfram.Client

//	функция отображения событий
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel { //	отображение для каждого события
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	godotenv.Load(".env")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN")) //	подключаем бота с нашими токенами
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))                                 //	подключаем модулю witai
	wolframClient := &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}                 //	подключаем модулю Вольфрам
	go printCommandEvents(bot.CommandEvents())                                           //	выводим сообщение о событии

	bot.Command("query for bot - <message>", &slacker.CommandDefinition{ //	бот будет отвечать на фразы записанные в виде "query for bot - <message>"
		Description: "send any question to wolfram",
		Example:     "who is the president of india",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("message") //	парсим запрос для получения сообщения

			msg, _ := client.Parse(&witai.MessageRequest{ //	парсим сообщение
				Query: query,
			})
			data, _ := json.MarshalIndent(msg, "", "    ") 	//	переводим в json с добавлением пробеллов
			rough := string(data[:]) 	//	переводим в строку
			value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value") 	//	вынимаем данные из запроса
			answer := value.String() 	//	переводим в строку
			res, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000) 	//	получаем ответ на наш запрос
			if err != nil {
				fmt.Println("there is an error") 	//	отлов ошибок
			}
			fmt.Println(value) 	//	отправляем ответ
			response.Reply(res)
		},
	})

	ctx, cancel := context.WithCancel(context.Background()) //	остановка работы бота
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil { //	отлов ошибок и логгирование
		log.Fatal(err)
	}
}
