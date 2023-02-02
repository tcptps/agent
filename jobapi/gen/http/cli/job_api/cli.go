// Code generated by goa v3.11.0, DO NOT EDIT.
//
// job_api HTTP client CLI support package
//
// Command:
// $ goa gen github.com/buildkite/agent/v3/jobapi/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	envc "github.com/buildkite/agent/v3/jobapi/gen/http/env/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `env get
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` env get` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		envFlags = flag.NewFlagSet("env", flag.ContinueOnError)

		envGetFlags = flag.NewFlagSet("get", flag.ExitOnError)
	)
	envFlags.Usage = envUsage
	envGetFlags.Usage = envGetUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "env":
			svcf = envFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "env":
			switch epn {
			case "get":
				epf = envGetFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "env":
			c := envc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get":
				endpoint = c.Get()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// envUsage displays the usage of the env command and its subcommands.
func envUsage() {
	fmt.Fprintf(os.Stderr, `Performs environment-related operations on the current job
Usage:
    %[1]s [globalflags] env COMMAND [flags]

COMMAND:
    get: Get implements get.

Additional help:
    %[1]s env COMMAND --help
`, os.Args[0])
}
func envGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] env get

Get implements get.

Example:
    %[1]s env get
`, os.Args[0])
}
