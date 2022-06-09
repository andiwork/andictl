package pack

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andiwork/andictl/utils"
	"github.com/spf13/viper"
	"golang.org/x/mod/modfile"
)

func CreatePackage() {
	modName := GetModuleName()
	tmplData := make(map[string]interface{})
	tmplData["modName"] = modName

	viper.SetConfigName("prod")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	port := viper.GetString("port")
	tmplData["port"] = port
	fmt.Println("Packaging ", modName, "application")
	tmpDir, err := ioutil.TempDir("", "andi.")
	defer os.Remove(tmpDir)
	if err != nil {
		fmt.Println("Unable to create tmp folder")
		os.Exit(0)
	}
	binaryFile := tmpDir + "/" + modName
	utils.ExecShellCommand("env", []string{"GOOS=linux", "GOARCH=amd64", "go", "build", "-o", binaryFile, "."}, false)
	utils.ExecShellCommand("cp", []string{"-r", "docs", tmpDir + "/docs"}, false)
	utils.ExecShellCommand("mkdir", []string{tmpDir + "/configs"}, false)
	utils.ExecShellCommand("cp", []string{"configs/prod.yaml", tmpDir + "/configs/app.yaml"}, false)

	tmpl, _ := packageTmpl.ReadFile("templates/systemd.service.gotmpl")
	utils.ProcessTmplFiles(tmpDir, modName+".service", tmpl, tmplData, false)

	tmpl, _ = packageTmpl.ReadFile("templates/nginx.conf.gotmpl")
	utils.ProcessTmplFiles(tmpDir, modName+".conf", tmpl, tmplData, false)

	utils.ExecShellCommand("tar", []string{"-czf", modName + ".tgz", "-C", tmpDir, "."}, false)
	fmt.Println("Packaging finished", tmpDir)
}
func GetModuleName() string {
	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		fmt.Printf("go.mod not found. You are not in Go project root folder.")
		os.Exit(0)
	}

	modName := modfile.ModulePath(goModBytes)
	return modName
}
