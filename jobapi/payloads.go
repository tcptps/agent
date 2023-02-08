package jobapi

// Error response is the response body for any errors that occur
type ErrorResponse struct {
	Error string `json:"error"`
}

// EnvGetResponse is the response body for the GET /env endpoint
type EnvGetResponse struct {
	Env map[string]string `json:"env"` // Different to EnvUpdateRequest because we don't want to send nulls
}

// EnvUpdateRequest is the request body for the GET /env endpoint
type EnvUpdateRequest struct {
	Env map[string]*string `json:"env"`
}

// EnvUpdateResponse is the response body for the PATCH /env endpoint
type EnvUpdateResponse struct {
	Added   []string `json:"added"`
	Updated []string `json:"updated"`
}

// EnvDeleteRequest is the request body for the DELETE /env endpoint
type EnvDeleteRequest struct {
	Keys []string `json:"keys"`
}

// EnvDeleteResponse is the response body for the DELETE /env endpoint
type EnvDeleteResponse struct {
	Deleted []string `json:"deleted"`
}
