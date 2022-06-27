package mublog

import "strings"

type Tag struct {
	tagFile  []rune
	articles []Article
}

var tagArticles map[string]Tag = make(map[string]Tag)

// Register article to belonging into certain tags
func RegisterTags(taglist string, tagDirname string, article Article) {
	tags := strings.Split(taglist, ", ")

	for i := range tags {
		if val, ok := tagArticles[tags[i]]; ok {
			val.articles = append(val.articles, article)
			tagArticles[tags[i]] = val
		} else {
			val := new(Tag)
			val.tagFile = []rune(tagDirname + "/" + strings.ToLower(tags[i]) + ".html")
			val.articles = append(val.articles, article)
			tagArticles[tags[i]] = *val
		}
	}
}

func CreateSortedTagLists(tagdir string) {

}
