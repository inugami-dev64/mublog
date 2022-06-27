package mublog

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// main configuration for mublog
// by default this file is located in /etc/mublog/mublog.conf
type Config struct {
	BlogMarkdownPath   string
	BlogHTMLPath       string
	TagHTMLPath        string
	IndexHTMLPath      string
	IndexTemplateHTML  string
	BlogTemplateHTML   string
	ListTemplateHTML   string
	MaxIndexedArticles int
	RSSFile            string
	RSSTitle           string
	RSSDescription     string
	RSSUrl             string
	Language           string
}

type Article struct {
	HTMLPath     []rune
	MarkdownPath []rune
	Tags         []rune
	Title        []rune
	PublishDate  []rune
	EditDate     []rune
}

type ArticleList []Article

func (a ArticleList) Less(i, j int) bool {
	return string(a[i].EditDate) > string(a[j].EditDate)
}

func (a ArticleList) Len() int {
	return len(a)
}

func (a ArticleList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func ReadConfiguration(confPath string) Config {
	_, err := os.Stat(confPath + "/mublog.conf")
	if err != err {
		log.Fatal("No configuration file found in: ", confPath+"/mublog.conf")
	}

	var config Config
	if _, err := toml.DecodeFile(confPath+"/mublog.conf", &config); err != nil {
		log.Fatal(err)
	}

	return config
}
