package handlers

import (
	"context"
	"net/http"

	"github.com/Ryan-Har/rharris.dev/src/templates/pages"
)

func Tools(w http.ResponseWriter, r *http.Request) {
	pages.Tools().Render(context.TODO(), w)
}
