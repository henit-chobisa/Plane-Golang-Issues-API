package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "planeapi",
	Short: "A Task project from plane.so, aiming for reducing latency from change in frameworks.",
	Long:  `The project is done as part of a Task from plane.so, which aims to recreate plane's issues api built using django in golang such that it targets to reduce the latency to under 500ms`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
