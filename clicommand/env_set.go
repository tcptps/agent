package clicommand

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/buildkite/agent/v3/bootstrap"
	"github.com/buildkite/agent/v3/env"
	"github.com/urfave/cli"
)

const envSetHelpDescription = `Usage:
  buildkite-agent env set [variable]

Description:
   Sets environment variable values in the current job execution environment. 
   Existing variables will be overwritten.

   Note that this subcommand is only available from within the job runner.
   
   Note that changes to the job environment variables only apply to subsequent
   phases of the job. To read the new values of variables from within the
   current phase, use env get.

   Note that Buildkite read-only variables cannot be overwritten.

Example (sets the variables LLAMA and ALPACA):

    $ buildkite-agent env set LLAMA=Kuzco "ALPACA=Geronimo the Incredible"
	
Example (sets the variables LLAMA and ALPACA using a JSON object over standard
input):

    $ buildkite-agent env set --format=json -
	{"ALPACA":"Geronimo the Incredible","LLAMA":"Kuzco"}`

type EnvSetConfig struct{}

var EnvSetCommand = cli.Command{
	Name:        "set",
	Usage:       "Sets variables in the job execution environment",
	Description: envSetHelpDescription,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "format",
			Usage:  "Input format: plain or json",
			EnvVar: "BUILDKITE_AGENT_ENV_SET_FORMAT",
			Value:  "plain",
		},
	},
	Action: envSetAction,
}

func envSetAction(c *cli.Context) error {
	cli, err := bootstrap.NewSocketClient()
	if err != nil {
		fmt.Fprintf(c.App.ErrWriter, "Could not create socket client: %v\nThis command can only be used from hooks or plugins running under the job runner.\n", err)
		os.Exit(1)
	}

	set := make(map[string]string)

	// Plain format
	parse := func(input string) error {
		e, v, ok := env.Split(input)
		if !ok {
			return fmt.Errorf("%q is not in key=value format", input)
		}
		set[e] = v
		return nil
	}

	// JSON format
	if c.String("format") == "json" {
		parse = func(input string) error {
			return json.Unmarshal([]byte(input), &set)
		}
	}

	// Inspect each arg, which could either be "-" or "KEY=value"
	for _, arg := range c.Args() {
		if arg == "-" {
			// Parse standard input
			sc := bufio.NewScanner(os.Stdin)
			line := 1
			for sc.Scan() {
				if err := parse(sc.Text()); err != nil {
					fmt.Fprintf(c.App.ErrWriter, "Couldn't parse input line %d: %v\n", line, err)
					os.Exit(1)
				}
				line++
			}
			if err := sc.Err(); err != nil {
				fmt.Fprintf(c.App.ErrWriter, "Couldn't scan the input buffer: %v\n", err)
				os.Exit(1)
			}
			continue
		}
		// Parse args directly
		if err := parse(arg); err != nil {
			fmt.Fprintf(c.App.ErrWriter, "Couldn't parse the command-line argument %q: %v\n", arg, err)
			os.Exit(1)
		}
	}

	// Encode to JSON
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(set); err != nil {
		fmt.Fprintf(c.App.ErrWriter, "Couldn't encode the environment variables into JSON: %v\n", err)
		os.Exit(1)
	}

	// Create the request
	req, err := http.NewRequestWithContext(context.Background(), "PATCH", "http://bootstrap/api/current-job/v0/env", &buf)
	if err != nil {
		fmt.Fprintf(c.App.ErrWriter, "Couldn't create a request: %v\n", err)
		os.Exit(1)
	}

	// Send the request
	resp, err := cli.Do(req)
	if err != nil {
		fmt.Fprintf(c.App.ErrWriter, "Couldn't perform the request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != 200 {
		fmt.Fprintln(c.App.ErrWriter, "The request failed:")
		io.Copy(c.App.ErrWriter, resp.Body)
		os.Exit(1)
	}

	// TODO: inspect the response for success
	io.Copy(c.App.Writer, resp.Body)

	return nil
}
