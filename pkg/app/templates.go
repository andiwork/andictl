package app

import "embed"

//go:embed templates/app.yaml.tmpl
var appTmpl embed.FS

//go:embed templates/prod.yaml.tmpl
var prodTmpl embed.FS
