package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

const confDesc = ``

var osType = ""

func newConfCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "conf [OS TYPE]",
		Short: "Initialization configuration OS",
		Long:  confDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("config %s", osType)
			return nil
		},
	}

	cmd.Flags().StringVarP(&osType, "type", "t", "linux", "...")
	return cmd
}
