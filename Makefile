proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=pkg --go_out=pkg api/proto/*.proto
imports:
	goimports -v -w -l ./
linter:
	golangci-lint run --enable-all -D gochecknoglobals --fix
deploy:
	kubectl apply -f ./deployments/deployments.yaml