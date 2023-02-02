package design

import . "goa.design/goa/v3/dsl"

var ErrorPayload = Type("ErrorPayload", func() {
	Attribute("error", String, "Error message")
	Example(Val{"error": "something went wrong"})
})

// API describes the global properties of the API server.
var _ = API("job_api", func() {
	Title("Job API")
	Description("An API exposed locally by the buildkite agent, to allow jobs to modify their own state")

	HTTP(func() {
		Path("/api/current-job/v0/")
		Consumes("application/json")
		Response("Internal server error", StatusInternalServerError, func() {
			Body(ErrorPayload)
		})

	})

	Server("job_api", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var EnvironmentPayload = Type("EnvironmentPayload", func() {
	Attribute("env", MapOf(String, String), "The environment variables for the current job")
	Example(Val{
		"env": map[string]any{
			"FOO":  "bar",
			"BAR":  nil,
			"PATH": "/usr/local/bin:/usr/bin:/bin",
		},
	})
})

var EnvUpdateResponse = Type("EnvUpdatePayload", func() {
	Attribute("added", ArrayOf(String), "The environment variables that were added")
	Attribute("removed", ArrayOf(String), "The environment variables that were removed")
	Attribute("changed", ArrayOf(String), "The environment variables that were changed")
	Example(Val{
		"added":   []string{"FOO"},
		"removed": []string{"BAR"},
		"changed": []string{"PATH"},
	})
})

// Service describes a service
var _ = Service("env", func() {
	Description("Performs environment-related operations on the current job")
	// Method describes a service method (endpoint)
	Method("get", func() {
		Result(EnvironmentPayload)
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			GET("/env")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})
	})
})
