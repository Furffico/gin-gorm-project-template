package config

// config entry
type Config struct {
	General  *GeneralConf  `mapstructure:"general"`
	Server   *ServerConf   `mapstructure:"server"`
	Database *DatabaseConf `mapstructure:"database"`
}

// config.general
type GeneralConf struct {
	Debug bool `mapstructure:"debug"`
}

// config.server
type ServerConf struct {
	BindAddress    string     `mapstructure:"bind_address"`
	BindPort       uint16     `mapstructure:"bind_port"`
	TrustedProxies []string   `mapstructure:"trusted_proxies"`
	APIBasePath    string     `mapstructure:"API_basepath"`
	Static         StaticConf `mapstructure:"static_file"`
}

// config.server.static_file
type StaticConf struct {
	Enabled   bool   `mapstructure:"enabled"`
	BasePath  string `mapstructure:"basepath"`
	Location  string `mapstructure:"location"`
	IndexPage string `mapstructure:"indexpage"`
}

// config.database
type DatabaseConf struct {
	Address  string `mapstructure:"address"`
	Port     uint16 `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Schema   string `mapstructure:"schema"`
}
