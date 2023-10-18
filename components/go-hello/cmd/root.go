package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version string

	rootCmd = &cobra.Command{
		Use:     "go-hello",
		Short:   "run go-hello operations",
		RunE: 	RunE,
		Version: version,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func RunE(cmd *cobra.Command, args []string) error {
	fmt.Println("Hello World")
	return nil
}
