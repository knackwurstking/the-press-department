<script>
  import { onMount } from "svelte";

  import { animate } from "./js/game";

  /** @type {HTMLCanvasElement} */
  let canvas;

  // some initial stuff
  let rbHz = 12;

  /** @type {import("./js/game").Game} */
  let game

  onMount(async () => {
    game = await animate(canvas);
  });
</script>

<div class="overlay">
  <input
    class="hz-rb"
    type="number"
    min={0}
    value={rbHz}
    on:change={(ev) => {
      for (let engine of game.engines) {
        engine.hz = parseInt(ev.currentTarget.value);
      }
    }}
  />
</div>

<main>
  <canvas bind:this={canvas} />
</main>

<style>
  main {
    width: 100vw;
    height: 100vh;
    display: flex;
    place-items: center;
    background-image: url("./Ground_248x248.png"),
      url("/assets/Ground_248x248.png");
  }

  canvas {
    touch-action: none;
  }

  .overlay {
    position: absolute;
    top: 0;
    right: 0;
    width: 4rem;
    height: 2.5rem;
    padding: 0.25rem;
  }

  .overlay input.hz-rb {
    z-index: 1;
    width: 100%;
    height: 100%;
    font-size: 1.5rem;
    text-align: center;
  }

  input[type="number"]::-webkit-outer-spin-button,
  input[type="number"]::-webkit-inner-spin-button {
    appearance: none;
    margin: 0;
  }

  input[type="number"] {
    appearance: textfield;
  }
</style>
