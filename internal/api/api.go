package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/dimiro1/tudu/internal/logging"
	"github.com/dimiro1/tudu/internal/storage"
	"github.com/dimiro1/tudu/internal/toolkit/render"
)

// GetSingleTodoHandler handle HTTP GET /todos/{id}.
type GetSingleTodoHandler struct {
	store  storage.Store
	logger *logging.Logger
	name   string
}

func NewGetSingleTodoHandler(store storage.Store, logger *logging.Logger) (*GetSingleTodoHandler, error) {
	if store == nil {
		return nil, errors.New("api: store cannot be nil")
	}

	if logger == nil {
		return nil, errors.New("api: logger cannot be nil")
	}

	return &GetSingleTodoHandler{store: store, logger: logger, name: "GetSingleTodoHandler"}, nil
}

// ServeHTTP fetch an item by id.
//
// Example:
// {
//     "id": 1,
//     "title": "Finish my homework",
//     "completed": false
//  }
// GET /todos/{id}
func (g *GetSingleTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	renderer, err := render.RendererFromRequest(r)
	if err != nil {
		g.logger.ErrorInvalidRenderer(err, g.name)
		return
	}

	var (
		pathVars = mux.Vars(r)
		strID    = pathVars["id"]
	)

	todoID, err := strconv.Atoi(strID)
	if err != nil {
		if encodeErr := renderer.Render(w, r, http.StatusBadRequest,
			errors.New("the given id is not valid")); encodeErr != nil {
			g.logger.ErrorRendering(encodeErr, g.name)
		}
		return
	}

	todo, err := g.store.Single(r.Context(), todoID)
	if err != nil {
		if encodeErr := renderer.Render(w, r, http.StatusInternalServerError, err); encodeErr != nil {
			g.logger.ErrorRendering(encodeErr, g.name)
		}
		return
	}

	if encodeErr := renderer.Render(w, r, http.StatusOK, todo); encodeErr != nil {
		g.logger.ErrorRendering(encodeErr, g.name)
	}
}

// GetTodosHandler handle HTTP GET /todos.
type GetTodosHandler struct {
	store  storage.Store
	logger *logging.Logger
	name   string
}

// NewGetTodosHandler returns a new GetTodosHandler handler.
func NewGetTodosHandler(store storage.Store, logger *logging.Logger) (*GetTodosHandler, error) {
	if store == nil {
		return nil, errors.New("api: store cannot be nil")
	}

	if logger == nil {
		return nil, errors.New("api: logger cannot be nil")
	}

	return &GetTodosHandler{store: store, logger: logger, name: "GetTodosHandler"}, nil
}

// ServeHTTP fetch all todos.
//
// Example:
// [
//     {
//         "id": 1,
//         "title": "Finish my homework",
//         "completed": false
//     },
//     {
//         "id": 2,
//         "title": "Develop a new sample app",
//         "completed": true
//     }
// ]
// GET /todos
func (g *GetTodosHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	renderer, err := render.RendererFromRequest(r)
	if err != nil {
		g.logger.ErrorInvalidRenderer(err, g.name)
		return
	}

	todos, err := g.store.All(r.Context())
	if err != nil {
		if encodeErr := renderer.Render(w, r, http.StatusInternalServerError, err); encodeErr != nil {
			g.logger.ErrorRendering(encodeErr, g.name)
		}
		return
	}

	if encodeErr := renderer.Render(w, r, http.StatusOK, todos); encodeErr != nil {
		g.logger.ErrorRendering(encodeErr, g.name)
	}
}
