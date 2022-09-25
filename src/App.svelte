<script>
  import { onMount, afterUpdate } from "svelte";

  import Data from "./js/data";


  /** @type {HTMLCanvasElement} */
  let canvas

  onMount(() => {
    console.log("on mount", canvas)
    // TODO: get the game running
  })
</script>

<main>
  <canvas bind:this={canvas}></canvas>
</main>

<style>
  main {
width: 100vw;
    height: 100vh;
    display: flex;
    place-items: center;
    justify-content: center;
  }

  canvas {
    border: var(--border);
    width: 1700px;
    height: 150px;
  }
</style>
