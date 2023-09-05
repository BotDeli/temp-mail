package spliter

import (
	"strings"
)

func SplitMail(mail string) (login, domain string) {
	sMail := strings.Split(mail, "@")
	return sMail[0], sMail[1]
}
