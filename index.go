package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/web"
	"github.com/mickelsonm/martini-dmvoauth/controllers/authentication"
	"github.com/mickelsonm/martini-dmvoauth/helpers/globals"
	"github.com/mickelsonm/martini-dmvoauth/helpers/oauth"
	"github.com/ottemo/dmv"
)

func main() {
	m := martini.Classic()
	m.Use(web.ContextWithCookieSecret(""))
	m.Use(sessions.Sessions(
		globals.SESSION_KEY, sessions.NewCookieStore([]byte(globals.SESSION_KEY))))
	m.Use(render.Renderer())

	m.Get("/", authentication.Base)
	m.Get("/logout", authentication.Logout)

	m.Get("/auth/google", dmv.AuthGoogle(oauth.GoogleOAuth2Options))
	m.Get("/oauth2callback", dmv.AuthGoogle(oauth.GoogleOAuth2Options), authentication.GoogleOAuth2Callback)

	m.Run()
}
