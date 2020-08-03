package api

import (
	"net/http"
)

// GetGotrueEnv is the endpoint for fetching current goauth email config
func (a *API) GetGotrueEnv(w http.ResponseWriter, r *http.Request) error {

	return sendJSON(w, http.StatusOK, 0)
}

// SetGotrueEnv is the endpoint for fetching current goauth email config
func (a *API) SetGotrueEnv(w http.ResponseWriter, r *http.Request) error {

	return sendJSON(w, http.StatusOK, 0)
}