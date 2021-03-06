root	:=		$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: clean build-linux build-mac build-win

clean:
	rm -f ContentSearcher
	rm -f ContentSearcher.app
	rm -f ContentSearcher.exe

build-linux:
	cd src && GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/ContentSearcher

build-mac:
	cd src && GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/ContentSearcher.app

build-win:
	cd src && GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/ContentSearcher.exe
