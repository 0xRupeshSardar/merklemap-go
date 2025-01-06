package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	// Define flags
	domain := flag.String("d", "", "Domain name to search for")
	outputFile := flag.String("o", "", "File to save the output (optional)")
	flag.Parse()

	// Validate domain input
	if *domain == "" {
		fmt.Println("Error: Domain name is required. Use -d flag to specify.")
		os.Exit(1)
	}

	// Define the output destination (file or terminal)
	var output io.Writer
	if *outputFile != "" {
		// Create or overwrite the output file
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
		fmt.Printf("Output will be saved to %s\n", *outputFile)
	} else {
		// Write to the terminal
		output = os.Stdout
	}

	baseURL := "https://api.merklemap.com/search"
	page := 0
	retryCount := 0
	const maxRetries = 5

	for {
		// Define query parameters
		params := url.Values{}
		params.Add("query", *domain)
		params.Add("page", fmt.Sprintf("%d", page))

		// Create the request
		req, err := http.NewRequest("GET", baseURL, nil)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			os.Exit(1)
		}
		req.URL.RawQuery = params.Encode()
		req.Header.Set("Accept", "application/json")

		// Make the HTTP request
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		// Check the HTTP status code
		if resp.StatusCode == http.StatusTooManyRequests {
			retryCount++
			if retryCount > maxRetries {
				fmt.Println("Error: Exceeded maximum retry attempts due to rate-limiting.")
				os.Exit(1)
			}
			fmt.Printf("Rate limit reached. Retrying in %d seconds...\n", 10*retryCount)
			time.Sleep(time.Duration(10*retryCount) * time.Second)
			continue
		} else if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error: Received HTTP status %d\n", resp.StatusCode)
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Response: %s\n", string(body))
			os.Exit(1)
		}

		// Reset retry count on successful request
		retryCount = 0

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
			os.Exit(1)
		}

		// Parse the JSON response
		var jsonResponse map[string]interface{}
		err = json.Unmarshal(body, &jsonResponse)
		if err != nil {
			fmt.Printf("Error parsing JSON response: %v\n", err)
			fmt.Printf("Raw Response: %s\n", string(body)) // Log raw response for debugging
			os.Exit(1)
		}

		// Write response to the output (file or terminal)
		fmt.Fprintf(output, "%s\n", string(body))
		
		// Check if results are empty or if there's no meaningful update
		results, ok := jsonResponse["results"].([]interface{})
		if !ok || len(results) == 0 {
			fmt.Println("No more results to fetch. Exiting pagination.")
			break
		}

		// Increment the page counter
		page++
		time.Sleep(2 * time.Second) // Add a delay between requests to avoid hitting rate limits
	}

	if *outputFile == "" {
		fmt.Println("Output displayed in the terminal.")
	}
}
