current_dir := $(dir $(mkfile_path))

default: build

clean:
	@echo "Cleaning up previous build..."
	@rm -f  $(current_dir)bin/*
	@rm -rf $(current_dir)go_binary_files
	@rm -rf $(current_dir)go_files

build: clean
	@echo "Building app..."
	@mkdir go_binary_files
	@mkdir go_files
	@go build -o bin/playgo main.go

test:
	@echo "Testing app..."
	@go test ./...
