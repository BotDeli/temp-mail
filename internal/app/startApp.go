package app

import (
	"log"
	"temp-mail/internal/server"
)

func StartApplication() {
	log.Fatal(server.StartServer())
}
