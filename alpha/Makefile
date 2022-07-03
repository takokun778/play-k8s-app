export
PORT := 8080

lint:
	@cp ../core/.golangci.yml . && \
	golangci-lint run --fix && \
	rm -rf .golangci.yml
dev:
	@air
