package render

import (
	"encoding/xml"
	"errors"
	"net/http"
	"reflect"
)

// XMLRenderer implementation of Renderer for XML responses.
type XMLRenderer struct{}

func (XMLRenderer) ContentType() string {
	return "text/xml"
}

func (xr XMLRenderer) Render(w http.ResponseWriter, r *http.Request, status int, toRender interface{}) error {
	if w == nil {
		return errors.New("render: http.ResponseWriter cannot be nil")
	}

	if r == nil {
		return errors.New("render: *http.Request cannot be nil")
	}

	switch toRenderType := toRender.(type) {
	case error:
		toRender = struct {
			XMLName xml.Name `xml:"error"`
			Message string   `xml:"message,attr"`
		}{
			Message: toRenderType.Error(),
		}
	default:
		switch reflect.TypeOf(toRender).Kind() {
		case reflect.Slice, reflect.Array:
			toRender = struct {
				XMLName  xml.Name    `xml:"List"`
				Children interface{} `xml:"Child"`
			}{
				Children: toRender,
			}
		}
	}

	marshaledXML, err := xml.Marshal(toRender)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", xr.ContentType())
	w.WriteHeader(status)
	_, err = w.Write(marshaledXML)
	return err
}
