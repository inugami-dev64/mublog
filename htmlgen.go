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

	file, err := os.Create(string(article.HTMLPath))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	file.Write([]byte(strContent))
	file.Close()
}

func WriteBlogListHTML(config Config, articles []Article) {
	content, err := os.ReadFile(config.ListTemplateHTML)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var list []string
	for i := range articles {
		entry := "<li><a href=\"" +
			string(articles[i].URL) +
			"\">" +
			string(articles[i].EditDate) +
			" - " +
			string(articles[i].Title) +
			"</a></li>\n"
		list = append(list, entry)
	}

	strContent := strings.ReplaceAll(string(content), "{content}", strings.Join(list, ""))
	strContent = strings.ReplaceAll(strContent, "{title}", "List of all blog articles")

	file, err := os.Create(config.IndexPathHTML + "/bloglist.html")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	file.Write([]byte(strContent))
	file.Close()
}

func WriteIndexHTML(config Config, sortedArticles []Article) {
	content, err := os.ReadFile(config.IndexTemplateHTML)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	strContent := WriteTags(config, string(content))

	// iterate through articles
	var list []string
	for i := 0; i < Min(len(sortedArticles), config.MaxIndexedArticles); i++ {
		entry := "<li><a href=\"" +
			string(sortedArticles[i].URL) +
			"\">" +
			string(sortedArticles[i].EditDate) +
			" - " +
			string(sortedArticles[i].Title) +
			"</a></li>\n"
		list = append(list, entry)
	}

	strContent = strings.ReplaceAll(strContent, "{articles}", strings.Join(list, ""))

	// check if {see-more} link should be made present
	if len(sortedArticles) > config.MaxIndexedArticles {
		entry := "<a href=\"" +
			config.IndexURL +
			"/bloglist.html\">More articles</a>\n"
		strContent = strings.ReplaceAll(strContent, "{see-more}", entry)
	} else {
		strContent = strings.ReplaceAll(strContent, "{see-more}", "")
	}

	file, err := os.Create(config.IndexPathHTML + "/index.html")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	file.Write([]byte(strContent))
	file.Close()
}
