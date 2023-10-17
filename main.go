package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/net/html"
)

func fetchURL(targetURL *url.URL) {
	resp, err := http.Get(targetURL.String())
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer closeResponseBody(resp.Body)

	dirName := filepath.Join(outputDir, targetURL.Host, filepath.FromSlash(targetURL.Path)) + ".html"
	if err := os.MkdirAll(dirName, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	filename := filepath.Join(dirName, "index.html")
	saveContent(filename, resp.Body)
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
	dirName := filepath.Join(outputDir, url.Host, filepath.FromSlash(url.Path)) + ".html"
	filename := filepath.Join(dirName, "index.html")

	// Check the modification time of the local file
	fileInfo, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist:", filename)
		return
	} else if err != nil {
		fmt.Println("Error accessing file info:", err)
		return
	}

	lastModified := fileInfo.ModTime()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer closeFile(file)

	doc, err := html.Parse(file)
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
	fmt.Println("last_fetch:", lastModified.Format(time.RFC3339))
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

var outputDir string

func main() {
	// Define the metadata flag
	metadataFlag := flag.Bool("metadata", false, "Display metadata")

	// Define the output directory flag
	flag.StringVar(&outputDir, "output", ".", "Directory to save fetched web pages")

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
