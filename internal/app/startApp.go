package app

import (
	"log"
	"temp-mail/internal/server"
	"temp-mail/internal/storage/connections"
	"temp-mail/internal/storage/mails"
)

func StartApplication() {
	sm := mails.InitStorageMails()
	sc := connections.InitStorageConnections()
	log.Fatal(server.StartServer(sm, sc))
}
