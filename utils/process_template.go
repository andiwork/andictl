package utils

import (
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/andiwork/andictl/configs"
	"github.com/google/uuid"
	"github.com/metal3d/go-slugify"
)

func ProcessTmplFiles(folder string, dstFileName string, tmplFile []byte, tmplData interface{}, debug bool) {
	tmpl := template.New("app-conf").Funcs(template.FuncMap{
		"uuidWithOutHyphen": UuidWithOutHyphen,
		"andictlSlugify":    AndictlSlugify,
	})
	tmpl, err := tmpl.Parse(string(tmplFile))
	if err != nil {
		log.Fatal("Error Parsing template: ", err)
		return
	}
	filePath := folder + "/" + dstFileName
	if debug {
		err = tmpl.Execute(os.Stderr, tmplData)
	} else {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal("Error creating file. ", err)
			return
		}

		err = tmpl.Execute(file, configs.AppConfs.App)

	}
	if err != nil {
		log.Fatal("Error executing template. ", filePath, err)
	}

}

func UuidWithOutHyphen() (value string) {
	uuidWithHyphen := uuid.New()
	value = strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return
}

func AndictlSlugify(text string) string {
	return strings.ToLower(strings.ReplaceAll(slugify.Marshal(text), "-", ""))
}
