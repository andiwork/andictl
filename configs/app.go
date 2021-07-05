package configs

type Application struct {
	Type         string // survey will match the question and field names
	Name         string // or you can tag fields to match a specific name
	Auth         string
	Port         string
	DatabaseType string `mapstructure:"database-type"`
}

type GeneralConfig struct {
	App Application `mapstructure:"application"`
}

var App GeneralConfig

func init() {
	App = GeneralConfig{}
}
