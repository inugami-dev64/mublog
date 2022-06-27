package mublog

import (
	"log"
	"os"
)

var rssState bool = false

// Write initial rss
func initialiseRSS(config Config, file *os.File) {
	rssInit := "<rss version=\"2.0\">\n" +
		"<channel>\n" +
		"<title>" + config.RssTitle + "</title>\n" +
		"<language>" + config.RssLanguage + "</language>\n" +
		"<atom:link href=\"" + config.RssURL + "\" rel=\"self\" type=\"application/rss+xml/\"/>\n"

	file.WriteString(rssInit)
}

func WriteToRSS(config Config, sortedArticles ArticleList) {
	file, err := os.OpenFile(config.RssFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	initialiseRSS(config, file)

	// Add rss items
	for i := range sortedArticles {
		md, err := os.ReadFile(string(sortedArticles[i].MarkdownPath))
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		md = ParseMetadata(md)
		rawHTML := GenerateRawHTML(md)
		rss := "<item>\n" +
			"<title>" + string(sortedArticles[i].Title) + "</title>\n" +
			"<pubDate>" + string(sortedArticles[i].EditDate) + "<pubdate>\n" +
			"<guid>" + string(sortedArticles[i].HTMLPath) + "<pubdate>\n" +
			"<description>" + rawHTML + "</description>\n" +
			"</item>\n"

		_, err = file.WriteString(rss)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	_, err = file.WriteString("</channel>\n</rss>\n")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	file.Close()
}
