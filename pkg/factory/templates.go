package factory

import "embed"

//go:embed templates/factory.go.gotmpl
var factoryTmpl embed.FS

//go:embed templates/hello.go.gotmpl
var helloTmpl embed.FS

//go:embed templates/world.go.gotmpl
var worldTmpl embed.FS

//go:embed templates/README.md.gotmpl
var readmeTmpl embed.FS
