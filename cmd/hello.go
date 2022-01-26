package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "omni",
	Short: "Omni is a tiny swiss-army knife.",
	Run: func(cmd *cobra.Command, args []string) {
		println("Hello, World!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
