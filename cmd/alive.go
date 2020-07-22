package cmd

import (
	"fmt"

	"github.com/BraspagDevelopers/bphc/lib"
	"github.com/spf13/cobra"
)

// aliveCmd represents the alive command
var aliveCmd = &cobra.Command{
	Use:     "alive",
	Aliases: []string{"a", "liveness"},
	Short:   "Check if the site is alive",
	Long: `Sends a GET HTTP request to a site in order to check its liveness.
	If the site returns a status code in the range 200-299, it will be considered alive.
	If the site returns any other status code, the check will fail.
	
	When the check succedes, it will be produce an exit code of 0. Any failure will produce a difference exit code. Additionaly, there will always be a message in STDOUT when the check fails.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		baseURL := args[0]
		err := lib.LivenessCheck(baseURL, livenessPathFlag, verboseFlag)
		handle(err)
		fmt.Printf("The site %s is alive.\n", baseURL)
	},
}

var livenessPathFlag string

func init() {
	rootCmd.AddCommand(aliveCmd)
	aliveCmd.Flags().StringVar(&livenessPathFlag, "path", "dfom.htm", "The path for the liveness endpoint")
}
