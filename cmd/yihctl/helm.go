package main

import (
	"github.com/spf13/cobra"
)

func newHelmCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "helm",
		Short: "asdf",

		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
