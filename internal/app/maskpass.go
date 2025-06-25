package app

import (
	u "net/url"
	"strings"
)

func maskPasswordInDSN(dsn string) string {
	url, err := u.Parse(dsn)
	if err != nil {
		return dsn
	}

	if url.User == nil {
		return dsn
	}

	username := url.User.Username()
	url.User = u.User(username)

	result := url.String()

	return strings.Replace(result, username+"@", username+":***@", 1)
}
