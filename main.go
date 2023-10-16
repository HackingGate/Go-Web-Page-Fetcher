package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"time"

	"golang.org/x/net/html"
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
}

func saveContent(filename string, content io.Reader) {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println("Error creating directory for asset:", err)
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file for asset:", err)
		return
	}
	defer closeFile(file)

	_, err = io.Copy(file, content)
	if err != nil {
		fmt.Println("Error saving content to file:", err)
		return
	}

	fmt.Println("Content saved to:", filename)
}

func closeResponseBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		fmt.Println("Error closing response body:", err)
	}
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
	}
}

func fetchAndExtractMetadata(url *url.URL) {
	resp, err := http.Get(url.String())
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer closeResponseBody(resp.Body)

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	// Extract metadata
	numLinks, numImages := extractMetadata(doc)

	// Display metadata
	fmt.Println("site:", url.String())
	fmt.Println("num_links:", numLinks)
	fmt.Println("images:", numImages)
	fmt.Println("last_fetch:", time.Now().Format(time.RFC3339))
}

func extractMetadata(n *html.Node) (numLinks int, numImages int) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			numLinks++
		case "img":
			numImages++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childLinks, childImages := extractMetadata(c)
		numLinks += childLinks
		numImages += childImages
	}
	return
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

	// Fetch the URLs
	if *metadataFlag {
		for _, parsedURL := range parsedURLs {
			fetchAndExtractMetadata(parsedURL)
		}
	} else {
		for _, parsedURL := range parsedURLs {
			fetchURL(parsedURL)
		}
	}
}
