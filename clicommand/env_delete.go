package clicommand

import (
	"fmt"
	"os"

	"github.com/buildkite/agent/v3/bootstrap"
	"github.com/urfave/cli"
)

const envDeleteHelpDescription = `Usage:
  buildkite-agent env delete [variables]

Description:
   Deletes environment variables from the current job execution environment. 

   Note that this subcommand is only available from within the job runner.
   
   Note that changes to the job environment variables only apply to subsequent
   phases of the job. To read the new values of variables from within the
   current phase, use env get.

   Note that Buildkite read-only variables cannot be deleted.

Example (deletes the variables LLAMA and ALPACA):

   $ buildkite-agent env delete LLAMA,ALPACA
	
Example (deletes the variables LLAMA and ALPACA with a JSON list supplied over
standard input):
    
   $ buildkite-agent env delete --format=json -
   ["LLAMA","ALPACA"]`

type EnvDeleteConfig struct{}

var EnvDeleteCommand = cli.Command{
	Name:        "delete",
	Usage:       "Deletes variables from the job execution environment",
	Description: envDeleteHelpDescription,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "format",
			Usage:  "Input format: plain or json",
			EnvVar: "BUILDKITE_AGENT_ENV_DELETE_FORMAT",
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
