package oauth

import (
	"github.com/ottemo/dmv"
)

var (
	GoogleOAuth2Options = &dmv.OAuth2Options{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:3000/oauth2callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}
)
