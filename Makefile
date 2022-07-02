all: build

clean:
	rm -rf builds icegen.zip

build:
	mkdir -p builds/linux builds/windows builds/darwin
	env GOOS=linux GOARCH=amd64 go build -o builds/linux/icegen main.go
	env GOOS=windows GOARCH=amd64 go build -o builds/windows/icegen main.go
	env GOOS=darwin GOARCH=amd64 go build -o builds/darwin/icegen main.go
	cd builds && zip -r ../icegen.zip * ../templates