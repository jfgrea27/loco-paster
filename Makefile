all: build

staticcheck:
	staticcheck ./...

fumpt:
	gofumpt -w .

vet:
	go vet ./...

lint: fumpt vet staticcheck

build_ui:
	(cd web && npm run build)

build_be: lint
	go build -o loco-paster cmd/main.go

build: build_ui build_be

test:
	go test ./...
run: build test
	go run cmd/main.go & (cd web && npm start)

docker: test lint build_ui
	GOOS=linux GOARCH=amd64 go build -o loco-paster  cmd/main.go

	docker build -t loco-paster .
