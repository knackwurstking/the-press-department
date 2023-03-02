# The Press Department

## WASM

### Serve

Install if needed

```bash
go install github.com/hajimehoshi/wasmserve@latest
```

Run locally

```bash
wasmserve ./cmd/the-press-department
```

### Build

Build the game

```bash
env GOOS=js GOARCH=wasm go build -o build/wasm/the-press-department.wasm ./cmd/the-press-department
```

Copy the `wasm_exec.js` binary

```bash
cp $(go env GOROOT)/misc/wasm/wasm_exec.js build/wasm/
```

Create the HTML file `build/wasm/index.html`

```html
<!DOCTYPE html>
<script src="wasm_exec.js"></script>
<script>
  if (!WebAssembly.instantiateStreaming) {
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
      const source = await (await resp).arrayBuffer();
      return await WebAssembly.instantiate(source, importObject);
    };
  }

  const go = new Go();
  WebAssembly.instantiateStreaming(
    fetch("the-press-department.wasm"),
    go.importObject
  ).then(result => {
    go.run(result.instance);
  });
</script>
```
