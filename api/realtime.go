package api

import (
	"net/http"
)

// GetRealtimeEnv is the endpoint for fetching current goauth email config
func (a *API) GetRealtimeEnv(w http.ResponseWriter, r *http.Request) error {

	return sendJSON(w, http.StatusOK, 0)
}

// SetRealtimeEnv is the endpoint for fetching current goauth email config
func (a *API) SetRealtimeEnv(w http.ResponseWriter, r *http.Request) error {

	return sendJSON(w, http.StatusOK, 0)
}