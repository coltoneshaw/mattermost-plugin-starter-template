package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	pluginapi "github.com/mattermost/mattermost/server/public/pluginapi"
)

type TestHandler struct {
	*ErrorHandler
	pluginAPI *pluginapi.Client
}

func NewTestHandler(router *mux.Router, api *pluginapi.Client) *TestHandler {
	handler := &TestHandler{
		ErrorHandler: &ErrorHandler{},
		pluginAPI:    api,
	}

	testRouter := router.PathPrefix("/test").Subrouter()

	testRouter.HandleFunc("", withContext(handler.postTest)).Methods(http.MethodPost)
	testRouter.HandleFunc("", withContext(handler.getTest)).Methods(http.MethodGet)

	return handler
}

type TestResponse struct {
	userID string
}

func (h *TestHandler) getTest(c *Context, w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-ID")

	testResponse := TestResponse{
		userID: userID,
	}

	fmt.Println("userID", userID)

	ReturnJSON(w, &testResponse, http.StatusOK)
}

func (h *TestHandler) postTest(c *Context, w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-ID")
	testResponse := TestResponse{
		userID: userID,
	}

	fmt.Println("userID", userID)

	ReturnJSON(w, &testResponse, http.StatusOK)
}
