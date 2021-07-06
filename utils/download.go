package utils

import (
	"io"
	"log"
	"net/http"
	"os"
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
	log.Println("Downloading :", url, " To:", filepath)
	out, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return
}
