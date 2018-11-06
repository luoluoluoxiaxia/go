package main

import (
	"os"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

//type Tuling struct {
//	ReqType int `json:"reqtype"`
//	Perception
//	UserInfo
//}
//
//type Perception struct {
//	InputText
//}
//
//type InputText struct {
//	Text string
//}
//
//type UserInfo struct {
//	ApiKey string
//	UserId string
//}

func main() {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:1087")
	//os.Setenv("HTTPS_PROXY", "https://127.0.0.1:1087")
	//c := http.Client{
	//	// Transport: &http.Transport{
	//	//  Proxy: http.ProxyURL(urlproxy),
	//	// },
	//}
	//if resp, err := c.Get("https://www.google.com"); err != nil {
	//	log.Fatalln(err)
	//} else {
	//	defer resp.Body.Close()
	//	body, _ := ioutil.ReadAll(resp.Body)
	//	fmt.Printf("%s\n", body)
	//}
	bot, err := tgbotapi.NewBotAPI("561163886:AAHcYyGqfoxlyTPZT62kG_FBbF31RKGJd4o")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		MsgHandle(bot,update)
		ImageHandle(bot,update)
	}

}

func MsgHandle(bot *tgbotapi.BotAPI,update tgbotapi.Update){

	if update.Message.Text == "" {
		return
	}

	if (update.Message.Text == "123") {
		del := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)

		bot.DeleteMessage(del)
	}

	if (update.Message.Text == "广告") {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "这是广告，快删了")

		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func ImageHandle(bot *tgbotapi.BotAPI,update tgbotapi.Update){

	if update.Message.Photo == nil {
		return
	}

	//换取链接
}