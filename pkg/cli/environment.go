package cli

import "github.com/spf13/pflag"

// EnvSetting describes all of the environment settings.
type EnvSetting struct {
	User       string
	configPath string
	Debug      bool
}

func (s EnvSetting) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&s.configPath, "config", "c", s.configPath, "path to yihctl configuration")
}

func New() *EnvSetting {
	env := EnvSetting{
		User:       "Vista",
		configPath: "~/.yihctl",
		Debug:      false,
	}

	return &env
}
