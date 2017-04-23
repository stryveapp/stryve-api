package controller

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
	Error map[string]string `json:"errors,omitempty"`
}

// NewHTTPSuccess returns a successful request
func NewHTTPSuccess(v interface{}) *HTTPSuccess {
	resp := new(HTTPSuccess)
	resp.Success = true
	resp.Data = v

	return resp
}

// NewHTTPError returns a failed request
func NewHTTPError(errors map[string]string) *HTTPError {
	resp := new(HTTPError)
	resp.Success = false
	resp.Error = errors

	return resp
}
