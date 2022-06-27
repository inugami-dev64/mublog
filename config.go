package mublog

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// main configuration for mublog
// by default this file is located in /etc/mublog/mublog.conf
type Config struct {
	BlogPathMarkdown string
	BlogPathHTML     string
	TagPathHTML      string
	IndexPathHTML    string

	IndexTemplateHTML  string
	BlogTemplateHTML   string
	ListTemplateHTML   string
	MaxIndexedArticles int

	IndexURL string
	BlogURL  string
	TagsURL  string
	RssURL   string

	RssFile        string
	RssTitle       string
	RssDescription string
	RssLanguage    string
}

type Article struct {
	HTMLPath     []rune
	URL          []rune
	MarkdownPath []rune
	Tags         []rune
	Title        []rune
	PublishDate  []rune
	EditDate     []rune
}

// Article sorting interface
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

// String sorting interface
type sortStrings []string

func (s sortStrings) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortStrings) Len() int {
	return len(s)
}

func (s sortStrings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func ReadConfiguration(configFile string) Config {
	_, err := os.Stat(configFile)
	if err != err {
		log.Fatal("No configuration file found in: ", configFile)
	}

	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
