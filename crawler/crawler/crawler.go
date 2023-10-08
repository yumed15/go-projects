package crawler

import (
	"fmt"
	"sync"
)

func Crawl() {

	sitesChannel := make(chan string)        // crawled sites
	crawledSitesChannel := make(chan string) // sites to be crawled
	toCrawlCountChannel := make(chan int)    // count of how many sites are left to be crawled

	siteToCrawl := "https://www.theguardian.com/uk"

	go func() {
		crawledSitesChannel <- siteToCrawl
	}()

	var wg sync.WaitGroup

	go processCrawledSites(sitesChannel, crawledSitesChannel, toCrawlCountChannel)
	go closeChannels(sitesChannel, crawledSitesChannel, toCrawlCountChannel)

	var numCrawlerThreads = 50
	for i := 0; i < numCrawlerThreads; i++ {
		wg.Add(1)
		go crawlSite(&wg, sitesChannel, crawledSitesChannel, toCrawlCountChannel)
	}

	wg.Wait()
}

func crawlSite(wg *sync.WaitGroup, sitesChannel chan string, crawledSitesChannel chan string, toCrawlCountChannel chan int) {

	for webpageURL := range sitesChannel {
		fmt.Printf("%v \n", webpageURL)
		scrapeSite(webpageURL, crawledSitesChannel)
		toCrawlCountChannel <- -1
	}

	wg.Done()
}

func processCrawledSites(sitesChannel chan string, crawledSitesChannel chan string, toCrawlCountChannel chan int) {
	foundUrls := make(map[string]bool)

	for cl := range crawledSitesChannel {
		if !foundUrls[cl] {
			foundUrls[cl] = true
			toCrawlCountChannel <- 1
			sitesChannel <- cl
		}
	}
}

func closeChannels(sitesChannel chan string, crawledSitesChannel chan string, toCrawlCountChannel chan int) {
	count := 0

	for c := range toCrawlCountChannel {
		count += c
		if count == 0 {
			close(sitesChannel)
			close(crawledSitesChannel)
			close(toCrawlCountChannel)
		}
	}
}
