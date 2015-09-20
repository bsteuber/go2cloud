package status

import "net/http"

// The subset of http status codes we actually use.
const (
	OK                  = http.StatusOK
	BadRequest          = http.StatusBadRequest
	InternalServerError = http.StatusInternalServerError
)
