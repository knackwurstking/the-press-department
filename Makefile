run:
	@go run -v ./cmd/the-press-department

build_wasm:
	@env GOOS=js GOARCH=wasm go build -o build/wasm/the-press-department.wasm ./cmd/the-press-department
	@cp `go env GOROOT`/misc/wasm/wasm_exec.js build/wasm/
	@echo "<!DOCTYPE html><script src=\"wasm_exec.js\"></script><script>if (!WebAssembly.instantiateStreaming) { WebAssembly.instantiateStreaming = async (resp, importObject) => { const source = await (await resp).arrayBuffer(); return await WebAssembly.instantiate(source, importObject);};} const go = new Go(); WebAssembly.instantiateStreaming(fetch(\"the-press-department.wasm\"), go.importObject).then(result => { go.run(result.instance); });</script>" > build/wasm/index.html
