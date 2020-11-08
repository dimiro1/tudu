package api

import (
	"net/http"

	"github.com/dimiro1/tudu/internal/toolkit/render"
)

// StatusUpProbe is an HTTP probe that always returns
// `status up`/200 OK.
func StatusUpProbe() http.Handler {
	renderer := render.JSONRenderer{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = renderer.Render(w, r, http.StatusOK, map[string]string{
			"status": "up",
		})
	})
}

// LiveProbe is an HTTP probe that should be used to
// notify that the application is live.
func LiveProbe() http.Handler {
	return StatusUpProbe()
}

// ReadyProbe is an HTTP probe that should be used to
// notify that the application is ready to receive requests.
func ReadyProbe() http.Handler {
	return StatusUpProbe()
}
