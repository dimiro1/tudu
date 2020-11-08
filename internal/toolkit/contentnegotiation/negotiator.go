package contentnegotiation

import (
	"errors"
	"net/http"

	"github.com/dimiro1/tudu/internal/toolkit/contentnegotiation/internal/httputil"
)

// NegotiateContentType returns the best content-type of the request.
func NegotiateContentType(r *http.Request, offers []string, defaultOffer string) (string, error) {
	if r == nil {
		return defaultOffer, errors.New("contentnegotiation: *http.Request cannot be nil")
	}

	return httputil.NegotiateContentType(r, offers, defaultOffer), nil
}
