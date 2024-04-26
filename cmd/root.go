package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "streaming",
		Short:         "A CLI to run a streaming platform",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.AddCommand(NewAPICmd())
	root.AddCommand(NewWorker())
	root.AddCommand(NewPlayoutServer())

	return root
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
