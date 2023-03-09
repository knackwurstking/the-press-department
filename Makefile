run:
	@go run -v ./cmd/the-press-department

build_wasm:
	@env GOOS=js GOARCH=wasm go build -o wasm/the-press-department.wasm ./cmd/the-press-department
	@cp ./wasm/the-press-department.wasm ./svelte/public/the-press-department.wasm
