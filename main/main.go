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

const configFile string = "/etc/mublog/mublog.conf"

func SortArticles() []mublog.Article {
	var sortedTitleLinks []mublog.Article
	for _, article := range blogMarkdownMap {
		sortedTitleLinks = append(sortedTitleLinks, article)
	}

	sort.Sort(mublog.ArticleList(sortedTitleLinks))
	return sortedTitleLinks
}

func main() {
	for {
		blogMarkdownMap = make(map[string]mublog.Article)
		mublog.InitTags()

		// read configuration file
		config := mublog.ReadConfiguration(configFile)

		// scan for files in markdown_path
		files, err := ioutil.ReadDir(config.BlogPathMarkdown)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			htmlFileName := config.BlogPathHTML + "/" + strings.ReplaceAll(file.Name(), ".md", ".html")
			markdownFileName := config.BlogPathMarkdown + "/" + file.Name()
			content, _ := ioutil.ReadFile(markdownFileName)
			md := mublog.ParseMetadata(content)
			rawHTML := mublog.GenerateRawHTML(md)

			article := new(mublog.Article)
			article.URL = []rune(config.BlogURL + "/" + strings.ReplaceAll(file.Name(), ".md", ".html"))
			article.EditDate = []rune(mublog.Metadata["edit-date"])
			article.Title = []rune(mublog.Metadata["title"])
			article.MarkdownPath = []rune(markdownFileName)
			article.HTMLPath = []rune(htmlFileName)
			article.PublishDate = []rune(mublog.Metadata["publish-date"])
			article.Tags = []rune(mublog.Metadata["tags"])
			blogMarkdownMap[file.Name()] = *article

			mublog.RegisterTags(config, *article)
			mublog.WriteTagHTMLs(config)
			mublog.WriteArticleHTML(config, *article, rawHTML)
		}

		articles := SortArticles()
		mublog.WriteBlogListHTML(config, articles)
		mublog.WriteIndexHTML(config, articles)
		mublog.WriteToRSS(config, articles)
		time.Sleep(10 * time.Second)
	}
}
