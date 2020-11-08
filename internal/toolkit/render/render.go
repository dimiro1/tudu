package render

import (
	"errors"
	"net/http"

	"github.com/dimiro1/tudu/internal/toolkit/contentnegotiation"
)

// Renderer implementations render the given data into response.
type Renderer interface {
	// Render renders the data into the http.ResponseWriter.
	Render(w http.ResponseWriter, r *http.Request, status int, toRender interface{}) error

	// ContentType returns the content-type associated with the renderer.
	ContentType() string
}

// RendererFromRequest performs content negotiation to get the proper renderer.
// Defaults to JSONRenderer.
func RendererFromRequest(r *http.Request) (Renderer, error) {
	if r == nil {
		return nil, errors.New("render: *http.Request cannot be nil")
	}

	defaultOffer := "application/json"
	offers := []string{
		"application/json",
		"text/json",
		"application/xml",
		"text/xml",
	}

	contentType, err := contentnegotiation.NegotiateContentType(r, offers, defaultOffer)
	if err != nil {
		return nil, err
	}

	switch contentType {
	case "application/xml", "text/xml":
		return XMLRenderer{}, nil
	case "application/json", "text/json":
		fallthrough
	default:
		return JSONRenderer{}, nil
	}
}
