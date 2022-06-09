package pack

import "embed"

//go:embed templates/*
var packageTmpl embed.FS
