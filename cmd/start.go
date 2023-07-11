package cmd

import (
	"github.com/spf13/cobra"
)

var port int
var host string

var start = &cobra.Command{
	Use:   "start",
	Short: "Command for starting the fiber server",
	Run: func(cmd *cobra.Command, args []string) {
		// todo start the api
	},
}

func init() {
	rootCmd.AddCommand(start)
	start.PersistentFlags().StringVarP(&host, "host", "u", "http://localhost/", "Current host for the api")
	start.PersistentFlags().IntVarP(&port, "port", "p", 3000, "Port on which the application starts")
}
