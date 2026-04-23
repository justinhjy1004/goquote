# schema generation for frontend
go run cmd/schema/main.go

# build wasm
GOOS=js GOARCH=wasm go build -o main.wasm cmd/wasm/main.go