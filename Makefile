build:
	@go build -o ./bin/gitIgnore cmd/gitignore/main.go

run:
	@go run cmd/gitignore/main.go

clean:
	@rm ./bin/gitIgnore && rm -d ./bin/

full: build
	@./bin/gitIgnore