package app

import "embed"

//go:embed templates/app.yaml.gotmpl
var appTmpl embed.FS

//go:embed templates/prod.yaml.gotmpl
var prodTmpl embed.FS

//go:embed templates/app.go.gotmpl
var appGoTmpl embed.FS

//go:embed templates/gorm.go.gotmpl
var gormGoTmpl embed.FS

//go:embed templates/restful.go.gotmpl
var restfulGoTmpl embed.FS

//go:embed templates/swagger.go.gotmpl
var swaggerGoTmpl embed.FS

//go:embed templates/main.go.gotmpl
var mainGoTmpl embed.FS
