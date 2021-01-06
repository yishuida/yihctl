package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const gitlabDesc = ``

var gitlabCfg string

func newGitlabCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gitlab",
		Short: "",
		Long:  gitlabDesc,
	}
	flags := cmd.PersistentFlags()
	addGitlabFlags(flags)
	cmd.AddCommand(newGitlabListTagsCmd())

	return cmd
}

func addGitlabFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&gitlabCfg, "config", "c", "", "gitlab configuration file")
}

func initConfig() {
	if gitlabCfg != "" {
		viper.SetConfigFile(gitlabCfg)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		cmdLogger.WithField("config path", viper.ConfigFileUsed()).Info("The configuration file was read successfully")
		/*if err := viper.Unmarshal(&cfg); err != nil {
			cmdLogger.Panic(err)
		}*/
	} else {
		cmdLogger.Error(err)
	}
}
