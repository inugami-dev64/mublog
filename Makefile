SOURCES=mublog/mublogd.go\
		config.go\
		htmlgen.go\
		metadata.go\
		rss.go\
		tags.go

mublogd: $(SOURCES)
	go build ./mublog/mublogd.go

.PHONY: install clean
install: mublogd
	cp -r mublogd /usr/local/bin
	mkdir /etc/mublog
	touch /etc/mublog/mublog.conf

clean:
	rm mublogd