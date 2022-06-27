package mublog

import (
	"log"
	"os"
	"sort"
	"strings"
)

type Tag struct {
	tagFile  []rune
	tagURL   []rune
	articles []Article
}

var tagArticles map[string]Tag = make(map[string]Tag)

// Register article to belonging into certain tags
func RegisterTags(config Config, article Article) {
	var tags []string
	if string(article.Tags) != "" {
		tags = strings.Split(string(article.Tags), ", ")
	}

	for i := range tags {
		if val, ok := tagArticles[tags[i]]; ok {
			val.articles = append(val.articles, article)
			tagArticles[tags[i]] = val
		} else {
			val := new(Tag)
			val.tagFile = []rune(config.TagPathHTML + "/" + strings.ToLower(tags[i]) + ".html")
			val.tagURL = []rune(config.TagsURL + "/" + strings.ToLower(tags[i]) + ".html")
			val.articles = append(val.articles, article)
			tagArticles[tags[i]] = *val
		}
	}
}

func WriteTags(config Config, templateHTML string) string {
	// read template
	var tagNames []string
	for tagName, _ := range tagArticles {
		tagNames = append(tagNames, tagName)
	}
	sort.Sort(sortStrings(tagNames))

	var content []string
	for i := range tagNames {
		entry := "<a href=\"" +
			string(tagArticles[tagNames[i]].tagURL) + "\">" +
			tagNames[i] +
			"</a>\n"
		content = append(content, entry)
	}

	return strings.ReplaceAll(templateHTML, "{tags}", strings.Join(content, ""))
}

func WriteTagHTMLs(config Config) {
	// read the template
	templateHTML, err := os.ReadFile(config.ListTemplateHTML)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for key, val := range tagArticles {
		html := strings.ReplaceAll(string(templateHTML), "{title}", "List of articles with tag "+key)

		sort.Sort(ArticleList(val.articles))
		var content []string
		for i := range val.articles {
			entry := "<li><a href=\"" +
				string(val.articles[i].URL) + "\">" +
				string(val.articles[i].EditDate) +
				" - " +
				string(val.articles[i].Title) +
				"</a></li>\n"
			content = append(content, entry)
		}

		html = strings.ReplaceAll(html, "{content}", strings.Join(content, ""))

		fname := config.TagPathHTML + "/" + strings.ToLower(key) + ".html"
		file, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		file.WriteString(html)
		file.Close()
	}
}
