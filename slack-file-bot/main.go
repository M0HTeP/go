package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

// основная функция бота
func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3398622637444-3420207474624-hsnHf9BSQRtYWD8Uf6Y6SjHx") // создаем среду с токеном бота
	os.Setenv("CHANNEL_ID", "C03BQJE7LKW")   // создаем среду с ID канала
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))   // создаем переменную, в которую передаем нового клиента с токеном
	channelArr := []string{os.Getenv("CHANNEL_ID")}  // получаем ID канала
	fileArr := []string{"TEST.pdf"}    //  храним в переменной файл в виде строк

	for i := 0; i < len(fileArr); i++ {     //  проходимся циклом по файлу
		params := slack.FileUploadParameters{    //  получаем параметры для загрузки
			Channels: channelArr,    //  канал
			File:     fileArr[i],    //  сам файл
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)    //  в случае ошибки выводим сообщение
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)    //  выводим сообщение с именем файла и ссылкой загрузки
	}
}
