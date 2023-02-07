package jobapi

type ErrorResponse struct {
	Error string `json:"error"`
}

type EnvUpdateRequest struct {
	Env map[string]*string `json:"env"`
}

type EnvGetResponse struct {
	Env map[string]string `json:"env"` // Different to EnvUpdateRequest because we don't want to send nulls
}

type EnvUpdateResponse struct {
	Added   []string `json:"added"`
	Updated []string `json:"updated"`
}

type EnvDeleteRequest struct {
	Keys []string `json:"keys"`
}

type EnvDeleteResponse struct {
	Deleted []string `json:"deleted"`
}
