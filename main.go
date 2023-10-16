package main

import (
	"flag"
	"fmt"
)

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
}
