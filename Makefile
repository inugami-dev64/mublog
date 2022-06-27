SOURCES=mublog/mublog.go\
		config.go\
		htmlgen.go\
		metadata.go\
		rss.go\
		tags.go

mublog: $(SOURCES)
	go build -o mugenblog ./mublog/mublog.go

.PHONY: install clean
install: mugenblog
	cp -r mugenblog /usr/local/bin
	mkdir /etc/mublog
	touch /etc/mublog/mublog.conf

clean:
	rm mugenblog