package mails

type StorageMails interface {
	GetMail(uuid string) (string, error)
	NewMail() (string, string, error)
}
