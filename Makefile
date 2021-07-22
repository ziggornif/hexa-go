COVERAGE_DIR=coverage
MAIN=main.go


audit:
	gosec ./...

lint:
	revive -config defaults.toml -formatter friendly ./...

test:
	rm -rf $(COVERAGE_DIR) && mkdir $(COVERAGE_DIR)
	go test  ./... -coverprofile=$(COVERAGE_DIR)/coverage.out
	go tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html
	go tool cover -func=$(COVERAGE_DIR)/coverage.out
test-junit:
	go test ./... -v 2>&1 | go-junit-report -set-exit-code > $(COVERAGE_DIR)/junit.xml

build:
	go build -ldflags "-w -s" -o hexa-go $(MAIN)

unused:
	go mod tidy
	
run:
	./hexa-go

