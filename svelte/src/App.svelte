<script lang="ts">
  // go wasm init
  if (!WebAssembly.instantiateStreaming) {
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
      const source = await (await resp).arrayBuffer();
      return await WebAssembly.instantiate(source, importObject);
    };
  }

  //@ts-expect-error
  const go = new Go();
  WebAssembly.instantiateStreaming(
    fetch("the-press-department.wasm"),
    go.importObject
  ).then(result => {
    go.run(result.instance);
  });
</script>
