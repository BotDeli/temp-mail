package gMessages

import (
	"fmt"
	"net/http"
	"temp-mail/pkg/decoder"
	"temp-mail/pkg/spliter"
)

type InputMessage struct {
	Id      int    `json:"id"`
	From    string `json:"from"`
	Subject string `json:"subject"`
}

func GetMessagesFromMail(mail string) ([]InputMessage, error) {
	url := getUrlGetMessages(mail)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	messages, err := decoder.DecodeJSONToArray[InputMessage](response)
	return messages, err
}

func getUrlGetMessages(mail string) string {
	login, domain := spliter.SplitMail(mail)
	return fmt.Sprintf(
		"https://www.1secmail.com/api/v1/?action=getMessages&login=%s&domain=%s",
		login,
		domain)
}
