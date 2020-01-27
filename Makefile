BIN_DIR = bin

all:
	go build -o $(BIN_DIR)/server main.go handles.go

# test:
# 	go test cmd/parser/*.go
