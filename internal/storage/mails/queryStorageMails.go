package mails

import (
	"temp-mail/internal/api/gMail"
	"temp-mail/pkg/generators"
	"time"
)

func (r *RedisMails) GetMail(uuid string) (string, error) {
	cmd := r.Client.Get(r.Ctx, uuid)
	return cmd.Result()
}

func (r *RedisMails) NewMail() (string, string, error) {
	uuid := generators.NewUUID()
	mail, err := gMail.GetTempMail(r.Dom)
	if err != nil {
		return "", "", err
	}

	cmd := r.Client.Set(r.Ctx, uuid, mail, 30*time.Minute)
	return uuid, mail, cmd.Err()
}
