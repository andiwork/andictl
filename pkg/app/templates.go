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

//go:embed templates/jwt.go.gotmpl
var jwtGoTmpl embed.FS

//go:embed templates/authz.go.gotmpl
var authzGoTmpl embed.FS

//go:embed templates/swagger_helper.go.gotmpl
var swaggerHelperGoTmpl embed.FS

//go:embed templates/gitignore.gotmpl
var gitignoreTmpl embed.FS

//go:embed templates/db_singleton.go.gotmpl
var dbSingletionTmpl embed.FS

//go:embed templates/custom_gorm.go.gotmpl
var customGormGoTmpl embed.FS

//go:embed templates/custom_restful.go.gotmpl
var customRestfulGoTmpl embed.FS
