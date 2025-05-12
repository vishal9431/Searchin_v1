package main

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
	"crawler/config"
	"crawler/database"
	"crawler/utils"
)

var (
	CACHE_URLS []string = config.SEED_URLS
	OLD_URLS   []string
	urlsMutex  sync.Mutex
)

var (
	errorColor   = color.New(color.FgRed).SprintFunc()
	successColor = color.New(color.FgGreen).SprintFunc()
	warningColor = color.New(color.FgYellow).SprintFunc()
	infoColor    = color.New(color.FgCyan).SprintFunc()
)

func main() {
	database.InitDB()
	if err := utils.LoadUrls(&OLD_URLS); err != nil {
		fmt.Println(errorColor("[ERROR]:"), "Failed to load old URLs:", err)
		return
	}
	fmt.Println(infoColor("[INFO]:"), "Old URLs loaded:", len(OLD_URLS))

	startCrawling()
	fmt.Println(successColor("[DONE]:"), "Crawling completed")
}

func startCrawling() {
	var wg sync.WaitGroup

	for {
		urlsMutex.Lock()
		if len(CACHE_URLS) == 0 {
			urlsMutex.Unlock()
			break
		}
		currentBatch := CACHE_URLS
		CACHE_URLS = []string{}
		urlsMutex.Unlock()

		for _, url := range currentBatch {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				handleURL(url)
			}(url)
		}

		wg.Wait()
	}
}

func handleURL(url string) {
	if !utils.IsUrl(url) {
		fmt.Println(warningColor("[WARNING]:"), "Invalid URL:", url)
		return
	}

	data, err := utils.GetData(url)
	if err != nil {
		fmt.Println(errorColor("[ERROR]:"), "Failed to crawl", url, "=>", err)
		return
	}

	if links, ok := data["links"].([]string); ok {
		extractAndQueueLinks(links)
	}

	title, ok1 := data["title"].(string)
	description, ok2 := data["description"].(string)

	if ok1 && ok2 {
		saveIfValid(url, title, description)
	}
}

func extractAndQueueLinks(links []string) {
	urlsMutex.Lock()
	defer urlsMutex.Unlock()

	for _, link := range links {
		if !utils.Contains(OLD_URLS, link) && !utils.Contains(CACHE_URLS, link) {
			CACHE_URLS = append(CACHE_URLS, link)
		}
	}
}

func saveIfValid(url, title, description string) {
	success, _ := utils.AddContentIfNotExists(url, title, description)
	if success {
		fmt.Println(successColor("[SUCCESS]:"), "Saved:", url)
	} else {
		fmt.Println(warningColor("[WARNING]:"), "Already exists:", url)
	}
}
