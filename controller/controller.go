package controller

import "github.com/go-pg/pg"

// Handler is the handler for route controller actions
type Handler struct {
	DB *pg.DB
}

// BaseResponse is the base HTTP response
type BaseResponse struct {
	Success bool `json:"success"`
}

// HTTPSuccess extends the base HTTP response for returning
// successful requests
type HTTPSuccess struct {
	BaseResponse
	Data interface{} `json:"data,omitempty"`
}

// HTTPError extends the base HTTP response for returning
// failed requests
type HTTPError struct {
	BaseResponse
	Error []string `json:"errors,omitempty"`
}

// NewHTTPSuccess returns a successful request
func NewHTTPSuccess(v interface{}) *HTTPSuccess {
	resp := new(HTTPSuccess)
	resp.Success = true
	resp.Data = v

	return resp
}

// NewHTTPError returns a failed request
func NewHTTPError(errors []string) *HTTPError {
	resp := new(HTTPError)
	resp.Success = false
	resp.Error = errors

	return resp
}
