package cmd

import (
	"fmt"

	"github.com/BraspagDevelopers/bphc/lib"
	"github.com/spf13/cobra"
)

// healthCmd represents the health command
var healthCmd = &cobra.Command{
	Use:     "healthy",
	Aliases: []string{"h", "health", "healthcheck"},
	Short:   "Check if the site is healthy",
	Long: `Sends a GET HTTP request to a site in order to check its health.
	If the site returns a status code in the range 200-299 and the body is in JSON format and the value of the property IsHealthy is true, the site is considered healhty.
	If not, the check will fail.
	
	When the check succedes, it will be produce an exit code of 0. Any failure will produce a difference exit code. Additionaly, there will always be a message in STDOUT when the check fails`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		baseURL := args[0]
		inclsucc, err := cmd.Flags().GetBool("all")
		handle(err)
		err, msg := lib.HealthCheck(baseURL, healthcheckPathFlag, verboseFlag, inclsucc)
		handle(err)
		fmt.Printf(msg, baseURL)
	},
}

var healthcheckPathFlag string

func init() {
	rootCmd.AddCommand(healthCmd)
	healthCmd.Flags().StringVar(&healthcheckPathFlag, "path", "/healthcheck", "The path for the healthcheck endpoint")
	healthCmd.Flags().BoolP("all", "a", false, "Include successes in the result message")
}
