package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/DEMON1A/Walker/pkg/strings"
	"github.com/DEMON1A/Walker/pkg/search"
	"github.com/DEMON1A/Walker/pkg/toml"
)

func main() {
	dirPtr := flag.String("dir", ".", "Directory path to list files from")
	strPtr := flag.String("search", "", "String to search for in those files")
	regexPtr := flag.String("regex", "", "Regex to search with in those files")
	maxPtr := flag.Int("max", 4, "Max length for strings to print")
	sensitivePtr := flag.Bool("sensitive", false, "Search with case sensitive/insensitive mode")
	scanPtr := flag.Bool("scan", false, "Scan all the found strings using a regex dataset")
	excludePtr := flag.String("exclude", "", "Rules ids you want to exclude from the scan")
	flag.Parse()

	// Get the absolute path of the directory
	absPath, err := filepath.Abs(*dirPtr)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	// Walk through the directory and list files
	err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error accessing path:", err)
			return err
		}

		// Check if the path is a file
		if !info.IsDir() {
			result, err := strings.ReadStringsFromFile(path, *maxPtr)
			if err != nil {
				log.Fatal(err)
			}

			if *scanPtr {
				// Load the configuration from the TOML file
				config, err := toml.LoadConfig("config/gitleaks.toml")
				if err != nil {
					log.Fatalf("Error loading config: %v", err)
				}

				// Precompile regex patterns from the configuration
				rules := make(map[string]search.RuleWithDescription)
				for _, rule := range config.Rules {
					re, err := regexp.Compile(rule.Regex)
					if err != nil {
						log.Printf("Invalid regex pattern in rule '%s': %v", rule.ID, err)
						continue
					}
					rules[rule.ID] = search.RuleWithDescription{
						Regex:       re,
						Description: rule.Description,
					}
				}

				// Search results with the precompiled regex patterns and their descriptions
				search.SearchWithRegexes(rules, result, strings.SplitString(*excludePtr), path)
			} else {
				if *strPtr != "" {
					search.SearchStringInResults(*strPtr, result, *sensitivePtr, path)
				} else if *regexPtr != "" {
					search.SearchRegexInResults(*regexPtr, result, *sensitivePtr)
				} else {
					log.Fatal("You must choose a search method")
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directory:", err)
	}
}