package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "calc",
	Short: "Small calculator for everyday use",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from hugo")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}