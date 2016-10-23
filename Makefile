all: bin/vkfornerds

bin/vkfornerds:
	@which gb > /dev/null || go get github.com/constabulary/gb/...
	gb build vkfornerds

install: bin/vkfornerds
	cp bin/vkfornerds $(GOPATH)/bin

clean:
	rm -rf ./bin/
	rm -rf ./pkg/

.PHONY: all install clean bin/vkfornerds
