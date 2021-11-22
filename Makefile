BINARY_NAME=ip-receive-server

build:
	cd cmd/server; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)

dev:
	cd cmd/server; go build -o $(BINARY_NAME) -v
	cd cmd/server; ./$(BINARY_NAME) -config=/Users/uzdz/work/golang/ip-receive-server/configs/config-dev.yaml