package handlers

import (
	"context"
	"net/http"

	"github.com/Ryan-Har/rharris.dev/src/templates/pages"
	"github.com/Ryan-Har/rharris.dev/src/templates/partials"
)

func Home(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(context.TODO(), w)
}

func HomeProfilePic(w http.ResponseWriter, r *http.Request) {
	partials.ProfilePic().Render(context.TODO(), w)
}
