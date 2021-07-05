package app

import (
	"log"
	"os"

	"github.com/andiwork/andictl/configs"
)

func Generate() {
	test := "test/"
	log.Println("les configs ======", configs.App)
	// create app folder structure
	// folder configs
	os.MkdirAll(test+"configs", os.ModePerm)
	// folder docs/swagger-ui
	os.MkdirAll(test+"docs/swagger-ui", os.ModePerm)
	// folder pkg
	os.MkdirAll(test+"pkg", os.ModePerm)
	// generate
}
