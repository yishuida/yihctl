package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yishuida/yihctl/internal/version"
	"io"
	"text/template"
)

const versionDesc = `
Show the version fo Yihctl

This will print a representation the version of Yihctl.
The output will look something like this:

version.BuildInfo{Version:"v0.0.1", GitCommit:"fe51cd1e31e6a202cba7dead9552a6d418ded79a", GitTreeState:"clean", GoVersion:"go1.13.10"}

- Version is the semantic version of the release.
- GitCommit is the SHA for the commit that this version was built from.
- GitTreeState is "clean" if there are no local code changes when this binary was
  built, and "dirty" if the binary was built from locally modified code.
- GoVersion is the version of Go that was used to compile Yihctl.

When using the --template flag the following properties are available to use in
the template:

- .Version contains the semantic version of Helm
- .GitCommit is the git commit
- .GitTreeState is the state of the git tree when Helm was built
- .GoVersion contains the version of Go that Helm was compiled with
`

type versionOptions struct {
	short    bool
	template string
}

func newVersionCmd(out io.Writer) *cobra.Command {
	o := &versionOptions{}

	cmd := &cobra.Command{
		Use:               "version",
		Short:             "print the client version information",
		Long:              versionDesc,
		ValidArgsFunction: noCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.run(out)
		},
	}
	f := cmd.Flags()
	f.BoolVar(&o.short, "short", false, "print the version number")
	f.StringVar(&o.template, "template", "", "template for version string format")
	f.BoolP("client", "c", true, "display client version information")
	f.MarkHidden("client")

	return cmd
}

func (o *versionOptions) run(out io.Writer) error {
	if o.template != "" {
		tt, err := template.New("_").Parse(o.template)
		if err != nil {
			return err
		}
		return tt.Execute(out, version.Get())
	}
	fmt.Fprintln(out, formatVersion(o.short))
	return nil
}

func formatVersion(short bool) string {
	v := version.Get()
	if short {
		if len(v.GitCommit) >= 7 {
			return fmt.Sprintf("%s+g%s", v.Version, v.GitCommit[:7])
		}
		return version.GetVersion()
	}
	return fmt.Sprintf("%#v", v)
}
