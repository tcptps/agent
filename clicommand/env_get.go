package clicommand

import (
	"fmt"
	"os"

	"github.com/buildkite/agent/v3/bootstrap"
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
	Action: func(c *cli.Context) error {
		// TODO: implement
		_, _, err := bootstrap.ConnectToSocket()
		if err != nil {
			fmt.Fprintf(c.App.ErrWriter, "Could not connect to control socket: %v\nThis command can only be used from hooks or plugins running under the job runner.\n", err)
			os.Exit(1)
		}

		return nil
	},
}
