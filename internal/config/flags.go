package config

import "github.com/gookit/config/v2"

type FlagConf struct {
	ConfigPath string
}

var define_flags = []string{
	"config",
}

func parseFlag(flags *config.Config) *FlagConf {
	return &FlagConf{
		ConfigPath: flags.String("config", "config/docker.toml"),
	}
}
