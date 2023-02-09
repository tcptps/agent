package jobapi

import (
	"encoding/json"
	"sort"
)

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

// UnmarshalJSON is a custom unmarshaler for EnvUpdateResponse
// It uses the normal logic for JSON unmarshalling, but also sorts the Added and Updated slices, largely for ease of testing
func (e *EnvUpdateResponse) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, e)
	if err != nil {
		return err
	}

	sort.Strings(e.Added)
	sort.Strings(e.Updated)

	return nil
}

// MarshalJSON is a custom marshaler for EnvUpdateResponse
// It uses the normal logic for JSON marshalling, but also sorts the Added and Updated slices, largely for ease of testing
func (e EnvUpdateResponse) MarshalJSON() ([]byte, error) {
	sort.Strings(e.Added)
	sort.Strings(e.Updated)

	return json.Marshal(e)
}

// EnvDeleteRequest is the request body for the DELETE /env endpoint
type EnvDeleteRequest struct {
	Keys []string `json:"keys"`
}

// UnmarshalJSON is a custom unmarshaler for EnvDeleteRequest
// It uses the normal logic for JSON unmarshalling, but also sorts the keys slice, largely for ease of testing
func (e *EnvDeleteRequest) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, e)
	if err != nil {
		return err
	}

	sort.Strings(e.Keys)

	return nil
}

// MarshalJSON is a custom marshaler for EnvDeleteRequest
// It uses the normal logic for JSON marshalling, but also sorts the keys slice, largely for ease of testing
func (e EnvDeleteRequest) MarshalJSON() ([]byte, error) {
	sort.Strings(e.Keys)
	return json.Marshal(e)
}

// EnvDeleteResponse is the response body for the DELETE /env endpoint
type EnvDeleteResponse struct {
	Deleted []string `json:"deleted"`
}

// UnmarshalJSON is a custom unmarshaler for EnvDeleteResponse
// It uses the normal logic for JSON unmarshalling, but also sorts the Deleted slice, largely for ease of testing
func (e *EnvDeleteResponse) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, e)
	if err != nil {
		return err
	}

	sort.Strings(e.Deleted)

	return nil
}

// MarshalJSON is a custom marshaler for EnvDeleteResponse
// It uses the normal logic for JSON marshalling, but also sorts the Deleted slice, largely for ease of testing
func (e EnvDeleteResponse) MarshalJSON() ([]byte, error) {
	sort.Strings(e.Deleted)
	return json.Marshal(e)
}
