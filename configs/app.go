package configs

import "github.com/google/uuid"

type Application struct {
	AppId        uuid.UUID
	Type         string // application type api, front
	Name         string // application name
	Auth         bool   // secure application api true|false
	AuthType     string // auth type jwt|oidc
	Port         string // application port
	DatabaseType string `mapstructure:"database-type"` // application database-type
}

type Model struct {
	Name    string
	Package string
}

type GeneralConfig struct {
	App    Application `mapstructure:"application"`
	Models []Model     `mapstructure:"models"`
}

var AppDir = "./"

var AppConfs GeneralConfig

func init() {
	AppConfs = GeneralConfig{}
}
