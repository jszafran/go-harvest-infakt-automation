test:
	go test ./...

build:
	go build && mv geninvoice /usr/local/bin/ && echo "Build done."
