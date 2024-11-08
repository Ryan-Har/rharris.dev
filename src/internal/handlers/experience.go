package handlers

import (
	"context"
	"net/http"

	"github.com/Ryan-Har/rharris.dev/src/templates/pages"
)

func Experience(w http.ResponseWriter, r *http.Request) {
	pages.Experience().Render(context.TODO(), w)
}
