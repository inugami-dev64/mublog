package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/inugami-dev64/mublog"
)

// key: markdown file name
// value: article struct
var blogMarkdownMap map[string]mublog.Article

func SortArticles() []mublog.Article {
	var sortedTitleLinks []mublog.Article
	for _, article := range blogMarkdownMap {
		sortedTitleLinks = append(sortedTitleLinks, article)
	}

	sort.Sort(mublog.ArticleList(sortedTitleLinks))
	return sortedTitleLinks
}

func main() {
	blogMarkdownMap = make(map[string]mublog.Article)

	for {
		// read configuration file
		config := mublog.ReadConfiguration(".")

		// scan for files in markdown_path
		files, err := ioutil.ReadDir(config.BlogMarkdownPath)
		if err != nil {
			log.Fatal(err)
		}

		relist := false
		for _, file := range files {
			if value, exists := blogMarkdownMap[file.Name()]; exists {
				content, _ := ioutil.ReadFile(file.Name())
				md := mublog.ParseMetadata(content)

				// if any of the header content was modified, rewrite blog article
				if mublog.Metadata["title"] != string(value.Title) || mublog.Metadata["publish-date"] != string(value.PublishDate) || mublog.Metadata["edit-date"] != string(value.EditDate) || mublog.Metadata["tags"] != string(value.Tags) {
					value.Title = []rune(mublog.Metadata["title"])
					value.PublishDate = []rune(mublog.Metadata["publish-date"])
					value.EditDate = []rune(mublog.Metadata["edit-date"])
					value.Tags = []rune(mublog.Metadata["tags"])
					rawHTML := mublog.GenerateRawHTML(md)
					mublog.WriteArticleHTML(config, value, rawHTML)
					relist = true
				}
			} else {
				htmlFileName := config.BlogHTMLPath + "/" + strings.ReplaceAll(file.Name(), ".md", ".html")
				markdownFileName := config.BlogMarkdownPath + "/" + file.Name()
				content, _ := ioutil.ReadFile(markdownFileName)
				md := mublog.ParseMetadata(content)
				rawHTML := mublog.GenerateRawHTML(md)

				article := new(mublog.Article)
				article.EditDate = []rune(mublog.Metadata["edit-date"])
				article.Title = []rune(mublog.Metadata["title"])
				article.MarkdownPath = []rune(markdownFileName)
				article.HTMLPath = []rune(htmlFileName)
				article.PublishDate = []rune(mublog.Metadata["publish-date"])
				article.Tags = []rune(mublog.Metadata["tags"])
				blogMarkdownMap[file.Name()] = *article

				mublog.RegisterTags(string(article.Tags), config.TagHTMLPath, *article)
				mublog.WriteArticleHTML(config, *article, rawHTML)
				relist = true
			}
		}

		if relist {
			articles := SortArticles()
			mublog.WriteToRSS(config, articles)
		}

		// sleep for 60 seconds
		time.Sleep(60 * time.Second)
	}
}
