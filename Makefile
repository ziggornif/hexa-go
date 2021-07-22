CLIENT_DIR=client
CLIENT_DIST_DIR=static
COVERAGE_DIR=coverage
MAIN=main.go


audit:
	gosec ./...

lint:
	revive -config defaults.toml -formatter friendly ./...

test:
	rm -rf $(COVERAGE_DIR) && mkdir $(COVERAGE_DIR)
	go test ./... -v 2>&1 | go-junit-report -set-exit-code > coverage/junit.xml

build:
	go build -ldflags "-w -s" -o hexa-go $(MAIN)

unused:
	go mod tidy
	
run:
	./hexa-go

