package authentication

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/web"
	"github.com/ottemo/dmv"
	"net/http"
)

type SessionData struct {
	UserID      string
	UserName    string
	UserEmail   string
	DisplayName string
}

func Base(ctx *web.Context, s sessions.Session, r render.Render) {
	if uid := s.Get("userID"); uid == nil {
		r.HTML(http.StatusOK, "login", nil)
		return
	}
	r.HTML(http.StatusOK, "home", SessionData{
		UserID:      s.Get("userID").(string),
		DisplayName: s.Get("userDisplayName").(string),
		UserEmail:   s.Get("userEmail").(string),
	})
}

func Logout(ctx *web.Context, s sessions.Session) {
	s.Clear()
	ctx.Redirect(http.StatusFound, "/")
}

func GoogleOAuth2Callback(goog *dmv.Google, r render.Render, s sessions.Session, ctx *web.Context) {
	if len(goog.Errors) > 0 {
		ctx.Abort(http.StatusInternalServerError, "OAuth failure")
		return
	}
	s.Set("userID", goog.Profile.ID)
	s.Set("userEmail", goog.Profile.Email)
	s.Set("userDisplayName", goog.Profile.DisplayName)

	r.HTML(http.StatusOK, "home", SessionData{
		UserID:      goog.Profile.ID,
		DisplayName: goog.Profile.DisplayName,
		UserEmail:   goog.Profile.Email,
	})
}
