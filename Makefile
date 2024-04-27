run: 
	@go run ./01-helloworld/hello.go

docs: 
	godoc -http :8081

test:
	grc go test ./... -v