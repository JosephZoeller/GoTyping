all: release_linux run_release_linux

build_linux:
	go build ./cmd/typetest
	go build ./cmd/typetestd

release_linux:
	mkdir -p ./Release_Build/linux-amd64/server/web
	go build -o ./Release_Build/linux-amd64/GoTyping ./cmd/typetest
	go build -o ./Release_Build/linux-amd64/server/GoTyping-Server ./cmd/typetestd
	cp ./sentences.txt ./Release_Build/linux-amd64/sentences.txt
	cp ./config.json ./Release_Build/linux-amd64/config.json
	cp ./web/tables.html ./Release_Build/linux-amd64/server/web/tables.html

release_windows:
	mkdir -p ./Release_Build/windows-amd64/server/web/
	env GOOS=windows GOARCH=amd64 go build -o ./Release_Build/windows-amd64/GoTyping.exe ./cmd/typetest
	env GOOS=windows GOARCH=amd64 go build -o ./Release_Build/windows-amd64/server/GoTyping-Server.exe ./cmd/typetestd
	cp ./sentences.txt ./Release_Build/windows-amd64/sentences.txt
	cp ./config.json ./Release_Build/windows-amd64/config.json
	cp ./web/tables.html ./Release_Build/windows-amd64/server/web/tables.html

run_release_linux:
	./Release_Build/linux-amd64/GoTyping

release: release_linux release_windows