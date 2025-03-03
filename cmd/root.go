package cmd

import (
	"fmt"
	"os"

	"github.com/astr0n8t/inotify-tasker/internal"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "inotify-tasker",
	Short: "Update files on a loop for programs that utilize inotifywait",
	Long: `Allows a way to update files locally so that local 
	filesystem watching applications can be notified`,
	// Simply call the internal run command
	Run: func(cmd *cobra.Command, args []string) {
		internal.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}
