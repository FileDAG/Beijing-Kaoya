package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadFromUrl(urlToDownload string) (string, error) {
	resp, _ := http.Get(urlToDownload)
	defer resp.Body.Close()

	segments := strings.Split(urlToDownload, "/")
	filename := segments[len(segments)-2]
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println("Finish downloading:", filename)
	return filename, nil
}
