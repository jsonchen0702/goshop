doc:
	swag fmt && swag init -d ./cmd/api

unittest:
	go test -timeout 9000s -a -v -coverprofile=coverage.out -coverpkg=./... ./... 2>&1 | tee report.out

mock:
	go generate ./...