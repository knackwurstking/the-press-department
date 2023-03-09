run:
	@go run -v ./cmd/the-press-department

build_wasm:
	@env GOOS=js GOARCH=wasm go build -o wasm/the-press-department.wasm ./cmd/the-press-department
	@cp `go env GOROOT`/misc/wasm/wasm_exec.js svelte/public/wasm_exec.js
