root	:=		$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: clean build-linux build-mac build-win

clean:
	rm -f ${root}/ContentSearcher
	rm -f ${root}/ContentSearcher.app
	rm -f ${root}/ContentSearcher.exe
	rm -f ${root}/webviewCli

build-linux:
	cd ${root}/src/server && GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/ContentSearcher
	cd ${root}/src/cli    && GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/webviewCli

build-mac:
	cd ${root}/src/server && GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/ContentSearcher.app
	cd ${root}/src/cli    && GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/webviewCli

build-win:
	cd ${root}/src/server && GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/ContentSearcher.exe
	cd ${root}/src/cli    && GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/webviewCli
