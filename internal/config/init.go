package config

import (
	"log"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

var (
	Conf  *Config
	Flags *FlagConf
)

// Load command-line arguments
func LoadFlags() (err error) {
	if Flags != nil {
		return
	}
	flags := config.NewEmpty("flags")
	err = flags.LoadFlags(define_flags)
	if err != nil {
		return
	}
	Flags = parseFlag(flags)
	return
}

// Load config file (toml format)
func LoadConfig() (err error) {
	if Conf != nil { // Config is required to load only once
		return
	}

	// Invoke LoadFlags before LoadConfig
	err = LoadFlags()
	if err != nil {
		return
	}

	config.WithOptions(config.ParseEnv)
	config.AddDriver(toml.Driver)

	confpath := Flags.ConfigPath
	err = config.LoadFiles(confpath)
	if err != nil {
		return
	}

	err = config.BindStruct("", &Conf)
	if err == nil {
		log.Printf("Configuration loaded from `%s`", confpath)
	}
	return
}
