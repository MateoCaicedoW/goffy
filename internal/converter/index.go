package converter

import (
	"net/http"

	"go.leapkit.dev/core/server"
)

// Renders the home page of the application.
func Index(w http.ResponseWriter, r *http.Request) {

	l := layout{
		Title:       "GoFFY - Home",
		Description: "Some description",
		Yield:       indexEl(),
	}.New()

	if err := l.Render(w); err != nil {
		server.Errorf(w, http.StatusInternalServerError, "error rendering layour %w", err)
		return
	}
}
