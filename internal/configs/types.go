package configs

type Config struct {
	Service  Service  `mapstructure:"service"`
	Database Database `mapstructure:"database"`
}

type Service struct {
	Port      string `mapstructure:"port"`
	SecretJWT string `mapstructure:"secretJWT"`
}

type Database struct {
	DataSourceName string `yaml:"dataSourceName"`
}
