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

If you want to embed your game into your web page, using iframe is strongly
recommended. The screen scale is automatically adjusted.
If the above HTML's name is main.html, the host HTML will be like this:

```html
<!DOCTYPE html>
<iframe src="main.html" width="640" height="480"></iframe>
```

You might find this message with Chrome:

The AudioContext was not allowed to start. It must be resume (or created)
after a user gesture on the page. [https://goo.gl/7K7WLu]

In this case, you can solve this by putting `allow="autoplay"` on the iframe.

## Mind Notes

### How does this `press`, `engines` and `tiles` thing work

- the `tiles` package contains info about the current product, tiles assets for
  each product, ...
- the `press` produces tiles at a given speed (ex.: 6 bumps per minute)
- the `engines` will transport the `tiles` from the `press` from A to B
  (right to left), engines can be configured (transport speed).
- in this first version, there is only one engine (just to simplify things)

- The `board` will turn on the `press` and the `engines`
  - just setup everything
- The `press` produces a tile and outputs the tile to the engine
  - this could be done with channels
- The engine will transport this tile from A to B (B is null)
