all: update build

update:
	@echo "updating..."
	go get

clear:
	@echo "cleaning up..."
	@rm icegen || true
	@rm icecast.xml || true

build:
	@echo "building..."
	go build
	@du -sh icegen
