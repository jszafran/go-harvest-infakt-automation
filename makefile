test:
	go test ./...

build:
	go build && echo "Build done." && mv geninvoice /usr/local/bin/ && echo "Finished."
