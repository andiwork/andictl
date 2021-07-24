package utils

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/metal3d/go-slugify"
)

func ProcessTmplFiles(folder string, dstFileName string, tmplFile []byte, tmplData interface{}, debug bool) {
	tmpl, err := template.New("andi").Funcs(template.FuncMap{
		"uuidWithOutHyphen": UuidWithOutHyphen,
		"andictlSlugify":    AndictlSlugify,
		"toLower":           strings.ToLower,
		"title":             strings.Title,
		"uuidNew":           uuid.New,
	}).Parse(string(tmplFile))
	tmpl = template.Must(tmpl, err)
	if err != nil {
		log.Fatal("Error Parsing template: ", err)
		return
	}
	filePath := folder + "/" + dstFileName
	if debug {
		err = tmpl.Execute(os.Stderr, tmplData)
	} else {
		file, err := os.Create(filePath)
		defer file.Close()
		if err != nil {
			log.Fatal("Error creating file. ", err)
			return
		}

		err = tmpl.Execute(file, tmplData)
		fmt.Println("create ", filePath)

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
