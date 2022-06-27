# Mublog static site generator

Mublog is a small blog site generator, written in Go. It allows to convert Markdown articles into HTML, 
create a list page for all articles sorted by publishing date and use tags for categorizing them.


## Getting started

First compile and install `mublogd` daemon.
```
# make install
```

Now in order to start using it, you will need to write a configuration file in `/etc/mublog/mublog.conf`.
Configuration file is quite straight forward and it should look like this:

```conf
# Path settings
BlogPathMarkdown = "/var/www/articles"
BlogPathHTML = "/var/www/myblog/blog"
TagPathHTML = "/var/www/myblog/tags"
IndexPathHTML = "/var/www/myblog"

# Templates
IndexTemplateHTML = "/var/www/templates/index.html"
BlogTemplateHTML = "/var/www/templates/blog.html"
ListTemplateHTML = "/var/www/templates/list.html"
MaxIndexedArticles = 5

# URL specifications
IndexURL = "https://example.org"
BlogURL = "https://example.org/blog"
TagsURL = "https://example.org/tags"
RssURL = "https://example.org/rss.xml"

# RSS specifications
RssFile = "/var/www/myblog/rss.xml"
RssTitle = "example.org"
RssDescription = "My blog's RSS feed"
RssLanguage = "en-US"
```

This configuration specifies directories to use for input / output data, templates for HTML generation 
and URLs that are going to be used. HTML templates use template identifiers with `{keyword}` syntax, 
which are used by generator to replace them with appropriate HTML code. 


| Identifier     | Description                                                           | Used by                            |
|----------------|-----------------------------------------------------------------------|------------------------------------|
| `{tags}`       | Specify the location of tag links (uses `<a>` elements without break) | IndexTemplateHTML                  |
| `{articles}`   | Specify the location of article links (uses `<li>` elements)          | IndexTemplateHTML                  |
| `{content}`    | General identifier to use for content                                 | BlogTemplateHTML, ListTemplateHTML |


Note that the `{content}` identifier can mean different thing for `BlogTemplateHTML` and `ListTemplateHTML`. For blog templates `{content}`
represents the location where generated blog contents should be, but for list templates this means the location where list elements will be put.

Once configuration steps are completed and templates are written, start the daemon
```
# mublogd
```