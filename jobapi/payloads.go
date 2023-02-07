package jobapi

type ErrorResponse struct {
	Error string `json:"error"`
}

type EnvPayload struct {
	Env map[string]string `json:"env"`
}

type EnvUpdateRequest EnvPayload
type EnvGetResponse EnvPayload

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
