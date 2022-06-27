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
	RSSPath            string
}

type Article struct {
	HTMLPath    []rune
	Tags        []rune
	Title       []rune
	PublishDate []rune
	EditDate    []rune
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
