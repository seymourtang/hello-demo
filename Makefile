build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -mod vendor -v -o ./bin/app ./cmd/main.go

container:
	@docker build -f ./Dockerfile -t seymourtang/hello-world:test .

test:
	@go test -v -cover $(go list ./...)
