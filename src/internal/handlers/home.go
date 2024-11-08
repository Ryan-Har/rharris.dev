package handlers

import (
	"context"
	"net/http"

	"github.com/Ryan-Har/rharris.dev/src/templates/pages"
)

func Home(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(context.TODO(), w)
}
