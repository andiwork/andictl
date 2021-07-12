package model

import "embed"

//go:embed templates/init.go.gotmpl
var initGoTmpl embed.FS

//go:embed templates/model.go.gotmpl
var modelGoTmpl embed.FS

//go:embed templates/model_resource.go.gotmpl
var modelResourceGoTmpl embed.FS

//go:embed templates/model_service.go.gotmpl
var modelServiceGoTmpl embed.FS

//go:embed templates/model_repository.go.gotmpl
var modelRepositoryGoTmpl embed.FS

//go:embed templates/gorm.go.gotmpl
var gormMigrateTmpl embed.FS

//go:embed templates/restful.go.gotmpl
var restfulWebserviceGoTmpl embed.FS
