package model

import "embed"

//go:embed templates/init.go.gotmpl
var initGoTmpl embed.FS

//go:embed templates/model.go.gotmpl
var modelGoTmpl embed.FS

//go:embed templates/model_resource.go.gotmpl
var modelResourceGoTmpl embed.FS
