package gMessages

import (
	"fmt"
	"net/http"
	"strings"
	"temp-mail/pkg/decoder"
)

type Message struct {
	Id      int    `json:"id"`
	From    string `json:"from"`
	Subject string `json:"subject"`
}

func GetMessagesFromMail(mail string) ([]Message, error) {
	url := getUrlGetMessages(mail)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	messages, err := decoder.DecodeJSON[Message](response)
	return messages, err
}

func getUrlGetMessages(mail string) string {
	splitMail := strings.Split(mail, "@")
	return fmt.Sprintf("https://www.1secmail.com/api/v1/?action=getMessages&login=%s&domain=%s", splitMail[0], splitMail[1])
}
