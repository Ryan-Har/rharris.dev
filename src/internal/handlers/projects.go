package handlers

import (
	"context"
	"net/http"

	"github.com/Ryan-Har/rharris.dev/src/templates/pages"
)

func Projects(w http.ResponseWriter, r *http.Request) {
	pages.Projects().Render(context.TODO(), w)
}
