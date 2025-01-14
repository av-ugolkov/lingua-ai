.PHONY: build
build:
	CGO_LDFLAGS="-L /lingua-ai/lib/x86_64-unknown-linux-gnu -lsherpa-onnx-c-api -lonnxruntime -Wl,-rpath,/lingua-ai/lib/x86_64-unknown-linux-gnu" go build --gcflags='all=-N -l' -o main cmd/main.go

.PHONY: run
run:
	@./run.sh run

.PHONY: release
release:
	@./deploy.sh release

.PHONY: dev
dev:
	@./deploy.sh dev

.PHONY: dev_ai
dev_ai:
	@./deploy.sh dev_ai
	
.PHONY: database.down
database.down:
	@echo ${shell ./deploy.sh database_down}

.PHONY: download_llm
download_llm:
	@echo ${shell ./download_llm.sh}

.PHONY: lint
lint:
	@go version
	@golangci-lint --version
	GOWORK=off golangci-lint run ./...

.PHONY: test
test: 
	go test ./... -count=1 -cover

.PHONY: gosec
gosec:
	gosec ./...

.PHONY: count-line
count-line:
	find . -name '*.go' | xargs wc -l

.PHONY: pprof.cpu
pprof.cpu:
	go tool pprof -http=:8080 profile

.PHONY: pprof.heap
pprof.heap:
	go tool pprof -http=:8080 heap

.PHONY: pprof.trace
pprof.trace:
	go tool pprof -http=:8080 trace
