package cmd

import (
	internal "github.com/ddubson/omni/internal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

func init() {
	rootCmd.AddCommand(NewWhoAmICmd())
}

func NewWhoAmICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "whoami",
		Short: "Print the user currently logged in.",
		Run: func(cmd *cobra.Command, args []string) {
			output, err := exec.Command("whoami").Output()
			if err != nil {
				log.Error("⚠️ Unable to fetch who you are 😢")
				os.Exit(1)
			}
			var user = strings.Replace(string(output), "\n", "", -1)
			internal.YellowPrint.Fprintf(cmd.OutOrStdout(), "⭐️ You are signed in as \""+user+"\"")
		},
	}
}
