package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

var list []string

func main() {
	fileUrl := "https://www.usom.gov.tr/url-list.txt"

	if err := DownloadFile("list.txt", fileUrl); err != nil {
		panic(err)
	}

	_, err := ioutil.ReadFile("list.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	file, err := os.Open("list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println()
		list = append(list, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Strings(list)

	fmt.Println(sort.SearchStrings(list, "abc-kemeja.fr"))
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
