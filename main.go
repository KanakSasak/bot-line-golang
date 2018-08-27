package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/gorilla/mux"
)

var bot *linebot.Client

func main() {
	var err error
	port := os.Getenv("PORT")
	//port = "9999"
	bot, err = linebot.New("<channel secret>", "<channel accsss token>")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/callback", callbackHandler)

	fmt.Println("Listen...")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

// IndexHandler is a representation of a index
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Wellcome to Gempa Line Bot - Get data from client"))

}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			//switch message := event.Message.(type) {
			//case *linebot.TextMessage:
			leftBtn := linebot.NewMessageAction("left", "left clicked")
			rightBtn := linebot.NewMessageAction("right", "right clicked")

			template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)

			message := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
			//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
			if _, err = bot.ReplyMessage(event.ReplyToken, message).Do(); err != nil {
				log.Print(err)
			}
			//}
		}
	}
}
