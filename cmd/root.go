package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(genCurveCmd)
	Log = log.New(os.Stderr, "", 0)
}

// Log is used by all the subcommands
var Log *log.Logger

var rootCmd = &cobra.Command{
	Use:   "dtcconfig",
	Short: "Creates DTC config files",
	Long: `Creates DTC Config Files.
	
	For more information, visit "https://github.com/niclabs/dtcconfig".`,
}

// Execute executes the program
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		Log.Printf("Error: %s", err)
		os.Exit(1)
	}
}
