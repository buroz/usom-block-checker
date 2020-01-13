package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var list []string

func main() {

	/*
		fileUrl := "https://www.usom.gov.tr/url-list.txt"

		if err := DownloadFile("list.txt", fileUrl); err != nil {
			panic(err)
		}
	*/

	data, err := ioutil.ReadFile("url-list.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	list = append(list, string(data))

}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
