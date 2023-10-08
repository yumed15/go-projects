package crawler

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"time"
)

func scrapeSite(webpageURL string, crawedLinksChannel chan string) {

	response, success := connectToSite(webpageURL)

	if !success {
		fmt.Println("Received error while connecting to website: ", webpageURL)
		return
	}

	defer response.Body.Close()

	tokenizer := html.NewTokenizer(response.Body)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			return
		}

		token := tokenizer.Token()

		if isAnchorTag(tokenType, token) {
			cl, ok := extractLinksFromToken(token, webpageURL)

			if ok {
				go func() {
					crawedLinksChannel <- cl
				}()
			}
		}
	}
}

func connectToSite(webpageURL string) (*http.Response, bool) {
	nilResponse := http.Response{}
	client := http.Client{
		Timeout: 60 * time.Second,
	}

	request, err := http.NewRequest("GET", webpageURL, nil)
	if err != nil {
		fmt.Println("Received error while creating new request: ", err)
		return &nilResponse, false
	}

	request.Header.Set("User-Agent", "GoBot v1.0 https://www.github.com/palvali/GoBot - This bot retrieves links and content.")

	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Received error while connecting to website: ", err)
		return &nilResponse, false
	}

	return response, true
}

func isAnchorTag(tokenType html.TokenType, token html.Token) bool {
	return tokenType == html.StartTagToken && token.DataAtom.String() == "a"
}

func extractLinksFromToken(token html.Token, webpageURL string) (string, bool) {
	for _, attr := range token.Attr {
		if attr.Key == "href" {
			link := attr.Val
			tl := formatURL(webpageURL, link)
			if tl == "" {
				break
			}
			return tl, true
		}
	}
	return "", false
}

func formatURL(base string, l string) string {

	base = strings.TrimSuffix(base, "/")

	switch {
	case strings.HasPrefix(l, "https://"):
	case strings.HasPrefix(l, "http://"):
		if strings.Contains(l, base) {
			return l
		}
		return ""
	case strings.HasPrefix(l, "/"):
		return base + l
	}
	return ""
}
