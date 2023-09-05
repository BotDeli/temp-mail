package gMessages

import (
	"fmt"
	"net/http"
	"temp-mail/pkg/decoder"
	"temp-mail/pkg/spliter"
)

type Message struct {
	Id      int    `json:"id"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Date    string `json:"date"`
	Body    string `json:"body"`
}

func ReadMessage(mail, id string) (Message, error) {
	url := getUrlReadMessage(mail, id)

	response, err := http.Get(url)
	if err != nil {
		return Message{}, err
	}

	msg, err := decoder.DecodeJSON[Message](response)
	return msg, err
}

func getUrlReadMessage(mail, id string) string {
	login, domain := spliter.SplitMail(mail)
	return fmt.Sprintf("https://www.1secmail.com/api/v1/?action=readMessage&login=%s&domain=%s&id=%s", login, domain, id)
}
