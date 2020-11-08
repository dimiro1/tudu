package render

import (
	"encoding/json"
	"errors"
	"net/http"
)

// JSONRenderer implementation of Renderer for JSON responses.
type JSONRenderer struct{}

func (JSONRenderer) ContentType() string {
	return "application/json"
}

func (jr JSONRenderer) Render(w http.ResponseWriter, r *http.Request, status int, toRender interface{}) error {
	if w == nil {
		return errors.New("render: http.ResponseWriter cannot be nil")
	}

	if r == nil {
		return errors.New("render: *http.Request cannot be nil")
	}

	switch toRenderType := toRender.(type) {
	case error:
		toRender = struct {
			Message string `json:"message"`
		}{
			toRenderType.Error(),
		}
	}

	marshaledJson, err := json.Marshal(toRender)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", jr.ContentType())
	w.WriteHeader(status)
	_, err = w.Write(marshaledJson)
	return err
}
