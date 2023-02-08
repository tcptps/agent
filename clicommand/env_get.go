package clicommand

import (
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

const envGetHelpDescription = `Usage:
  buildkite-agent env get [variables]

Description:
   Retrieves environment variables and their current values from the current job
   execution environment. 

   Note that this subcommand is only available from within the job runner.
   
   Changes to the job environment only apply to the environments of subsequent
   phases of the job. However, env get can be used to inspect the changes made
   with env set and env delete.

Example (gets the variables LLAMA and ALPACA):

    $ buildkite-agent env get LLAMA,ALPACA
	LLAMA=Kuzco
	ALPACA=Geronimo
	
Example (gets all variables):

    $ buildkite-agent env get --format=json-pretty
	{
		"ALPACA": "Geronimo",
		"LLAMA": "Kuzco"
	}`

type EnvGetConfig struct{}

var EnvGetCommand = cli.Command{
	Name:        "get",
	Usage:       "Gets variables from the job execution environment",
	Description: envGetHelpDescription,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "format",
			Usage:  "Output format: plain, json, or json-pretty",
			EnvVar: "BUILDKITE_AGENT_ENV_GET_FORMAT",
			Value:  "plain",
		},
	},
	Action: envGetAction,
}

func envGetAction(c *cli.Context) error {
	cli, err := bootstrap.NewSocketClient()
	if err != nil {
		fmt.Fprintf(c.App.ErrWriter, "Could not create socket client: %v\nThis command can only be used from hooks or plugins running under the job runner.\n", err)
		os.Exit(1)
	}

	req, err := http.NewRequestWithContext(context.Background(), "GET", "http://bootstrap/api/current-job/v0/env", nil)
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

	switch c.String("format") {
	case "json":
		// If it's json, output it directly.
		if _, err := io.Copy(c.App.Writer, resp.Body); err != nil {
			fmt.Fprintf(c.App.ErrWriter, "Couldn't read the response body: %v\n", err)
			os.Exit(1)
		}

	case "json-pretty":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(c.App.ErrWriter, "Couldn't read the response body: %v\n", err)
			os.Exit(1)
		}
		var buf bytes.Buffer
		if err := json.Indent(&buf, b, "", "  "); err != nil {
			fmt.Fprintf(c.App.ErrWriter, "Couldn't indent the output: %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(c.App.Writer, &buf); err != nil {
			fmt.Fprintf(c.App.ErrWriter, "Couldn't read the response body: %v\n", err)
			os.Exit(1)
		}

	case "plain":
		var vars env.Environment
		if err := json.NewDecoder(resp.Body).Decode(&vars); err != nil {
			fmt.Fprintf(c.App.ErrWriter, "Couldn't decode the response body: %v\n", err)
			os.Exit(1)
		}
		for _, v := range vars.ToSlice() {
			fmt.Println(v)
		}
	}

	return nil
}
