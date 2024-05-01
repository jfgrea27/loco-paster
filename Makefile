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
run: test
	go run cmd/main.go & (cd web && npm start)

docker: test lint build_ui
	GOOS=linux GOARCH=amd64 go build -o loco-paster  cmd/main.go

	docker build -t loco-paster .

deploy: docker
	docker tag loco-paster:latest ${AWS_ECR_REPO}/loco-paster:latest
	aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${AWS_ECR_REPO}
	#docker push ${AWS_ECR_REPO}/loco-paster:latest
