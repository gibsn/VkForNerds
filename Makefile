all: bin/vkfornerds

bin/vkfornerds:
	@which gb > /dev/null || go get github.com/constabulary/gb/...
	gb build vkfornerds

clean:
	rm -rf ./bin/
	rm -rf ./pkg/

tags:
	gotags -R ./ > tags

.PHONY: all install clean tags bin/vkfornerds
