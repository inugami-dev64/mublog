package mublog

import (
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

// Generate raw html data from markdown
func GenerateRawHTML(rawMD []byte) string {
	extensions := parser.CommonExtensions | parser.Titleblock ^ parser.DefinitionLists
	parser := parser.NewWithExtensions(extensions)

	html := []string{
		"<h1>" + Metadata["title"] + "</h1>\n",
		"<b class=\"date\">Publish date: " + Metadata["publish-date"] + "</b><br>\n",
		"<b class=\"date\">Last edited: " + Metadata["edit-date"] + "</b><br>\n",
		string(markdown.ToHTML(rawMD, parser, nil)),
	}

	return strings.Join(html, "")
}

func WriteArticleHTML(config Config, article Article, rawHTML string) {
	// read template HTML file
	content, err := os.ReadFile(config.BlogTemplateHTML)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	strContent := string(content)
	strContent = strings.ReplaceAll(strContent, "{title}", string(article.Title))
	strContent = strings.ReplaceAll(strContent, "{edit-date}", string(article.EditDate))
	strContent = strings.ReplaceAll(strContent, "{publish-date}", string(article.PublishDate))
	strContent = strings.ReplaceAll(strContent, "{content}", rawHTML)

	os.WriteFile(string(article.HTMLPath), []byte(strContent), 0440)
}
