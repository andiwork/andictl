package configs

type Application struct {
	Type         string // survey will match the question and field names
	Name         string // or you can tag fields to match a specific name
	Auth         bool
	AuthType     string
	Port         string
	DatabaseType string `mapstructure:"database-type"`
}

type Model struct {
	Name    string
	Package string
}

type GeneralConfig struct {
	App    Application `mapstructure:"application"`
	Models []Model     `mapstructure:"models"`
}

var AppDir = "test/"

var AppConfs GeneralConfig

func init() {
	AppConfs = GeneralConfig{}
}
