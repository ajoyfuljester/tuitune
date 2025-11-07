package cmd

import (
	"github.com/spf13/cobra"
)


var rootCmd = cobra.Command{
	Use: "main",
}

func Execute() {
	rootCmd.AddCommand(&cmdSearch)
	rootCmd.Execute()
}
