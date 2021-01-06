package main

import "github.com/spf13/cobra"

const GitlabListTagsDesc = `
`

func newGitlabListTagsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-tags",
		Short: "",
		Long:  GitlabListTagsDesc,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}
