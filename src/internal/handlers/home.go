package handlers

import (
	"context"
	"net/http"

	"github.com/Ryan-Har/rharris.dev/src/templates/layouts"
)

func Home(w http.ResponseWriter, r *http.Request) {
	layouts.Home().Render(context.TODO(), w)
}
