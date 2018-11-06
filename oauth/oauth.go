package oauth

import (
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

func init() {
	sessions.NewSession()
	oauth2.NewClient()
}
