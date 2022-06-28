SOURCES=main/main.go\
		config.go\
		htmlgen.go\
		metadata.go\
		rss.go\
		tags.go

mublogd: $(SOURCES)
	go build -o mublogd main/main.go

.PHONY: uninstall install clean

uninstall: 
	rm /usr/local/bin/mublogd

install: mublogd
	cp -r mublogd /usr/local/bin

clean:
	rm mugenblog