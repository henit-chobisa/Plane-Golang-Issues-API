package cmd

import (
	"fmt"

	"github.com/henit-chobisa/Plane-Golang-Issues-API/Packages/api"
	"github.com/henit-chobisa/Plane-Golang-Issues-API/Packages/constants"
	"github.com/spf13/cobra"
)

var port int
var host string

var start = &cobra.Command{
	Use:   "start",
	Short: "Command for starting the fiber server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(constants.Blue+"\nStarting Fiber API with host %v and port %v\n", host, port)
		api.Start(host, port)
	},
}

func init() {
	rootCmd.AddCommand(start)
	start.PersistentFlags().StringVarP(&host, "host", "u", "0.0.0.0", "Current host for the api")
	start.PersistentFlags().IntVarP(&port, "port", "p", 3000, "Port on which the application starts")
}
