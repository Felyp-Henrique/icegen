all: update build

update:
	@echo "updating..."
	go get

clear:
	@echo "cleaning up..."
	@rm icegen || true

build:
	@echo "building..."
	go build
	@du -sh icegen
