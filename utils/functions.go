package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(filename string, url string) (filepath string, err error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Create the file
	filepath = os.TempDir() + filename
	fmt.Println("Downloading :", url, " To:", filepath)
	out, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return
}

func InsertInfile(str string, placeHolder string, file string) {
	input, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, placeHolder) {
			lines[i] = "\t" + str + "\n" + placeHolder
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(file, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func PackageList(path string) (pack []string) {
	pack = []string{"new package"}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		if file.IsDir() && name != "middleware" {
			pack = append(pack, file.Name())
		}
	}
	return
}

func GetAppModule(goModPath string) string {
	file, err := os.Open(goModPath)
	if err != nil {
		fmt.Println("Generation can only be done at the root level of the application.", err)
		os.Exit(0)
	}
	r := bufio.NewReader(file)
	line, _, err := r.ReadLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	moduleName := bytes.TrimPrefix(line, []byte("module "))
	return string(moduleName)
}
