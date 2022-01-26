package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

func init() {
	rootCmd.AddCommand(whoAmICmd)
}

var whoAmICmd = &cobra.Command{
	Use:   "whoami",
	Short: "Print the user currently logged in.",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := exec.Command("whoami").Output()
		if err != nil {
			println("⚠️ Unable to fetch who you are 😢")
			os.Exit(1)
		}
		var user = strings.Replace(string(output), "\n", "", -1)
		println("⭐️ You are signed in as \"" + user + "\"")
	},
}
