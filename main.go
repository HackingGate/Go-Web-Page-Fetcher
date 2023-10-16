package main

import (
	"fetch/helper"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

func fetchURL(urlString string) {
	// Parse the URL
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)

	// Create a directory to store the file
	dirName := parsedURL.Host + parsedURL.Path + ".html"
	if err := os.MkdirAll(dirName, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	filename := path.Join(dirName, "index.html")
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error saving content to file:", err)
		return
	}

	fmt.Println("Content saved to:", filename)
}

func main() {
	// Define the metadata flag
	metadataFlag := flag.Bool("metadata", false, "Display metadata")

	// Parse the command line arguments
	flag.Parse()

	// Remaining arguments after flags are parsed
	args := flag.Args()

	// Check if URL is provided
	if len(args) < 1 {
		fmt.Println("Please provide a URL.")
		return
	}
	urls := args

	for _, url := range urls {
		if !helper.IsValidURL(url) {
			fmt.Println("Please provide a valid URL.")
			return
		}
	}

	// Display the provided URL
	fmt.Println("URLs:", urls)

	// Check if metadata flag is set
	if *metadataFlag {
		// Display metadata
		fmt.Println("site:")
		fmt.Println("num_links:")
		fmt.Println("images:")
		fmt.Println("last_fetch:")
	}

	// Fetch the URLs
	for _, urlString := range urls {
		fmt.Println("Fetching", urlString)
		fetchURL(urlString)
	}
}
