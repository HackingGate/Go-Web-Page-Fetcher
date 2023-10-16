package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

func fetchURL(url *url.URL) {
	resp, err := http.Get(url.String())
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
	dirName := url.Host + url.Path + ".html"
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
	urlStrings := args

	parsedURLs := make([]*url.URL, 0, len(urlStrings))

	for _, urlString := range urlStrings {
		parseURL, err := url.Parse(urlString)
		if err != nil {
			fmt.Println("Error parsing URL:", err)
			return
		}

		parsedURLs = append(parsedURLs, parseURL)
	}

	// Check if metadata flag is set
	if *metadataFlag {
		// Display metadata
		fmt.Println("site:")
		fmt.Println("num_links:")
		fmt.Println("images:")
		fmt.Println("last_fetch:")
	}

	// Fetch the URLs
	for _, parsedURL := range parsedURLs {
		fetchURL(parsedURL)
	}
}
