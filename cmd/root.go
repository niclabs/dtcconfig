package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(rsaCmd)
	rootCmd.AddCommand(genCurveCmd)
	Log = log.New(os.Stderr, "", 0)
}

var Log *log.Logger

var rootCmd = &cobra.Command{
	Use:   "dtcconfig",
	Short: "Creates DTC config files",
	Long: `Creates DTC Config Files.
	
	For more information, visit "https://github.com/niclabs/dtcconfig".`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		Log.Printf("Error: %s", err)
		os.Exit(1)
	}
}
