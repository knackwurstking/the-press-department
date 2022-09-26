<script>
  import { onMount } from "svelte";

  import Data from "./js/data";
  import Game from "./js/game";

  /** @type {HTMLCanvasElement} */
  let canvas;

  // some initial stuff
  let rbHz = 12;

  /** @type {Game} */
  let game;

  onMount(() => {
    const ctx = canvas.getContext("2d");
    canvas.width = 3460;
    canvas.height = 312;

    game = new Game(canvas, ctx, canvas.width, canvas.height, rbHz);

    // loading assets before runninng the game loop
    const queue = new Set();
    for (let asset of Data.assets) {
      game.assets[asset.name] = new Image(asset.width, asset.height);
      game.assets[asset.name].src = asset.src;

      queue.add(game.assets[asset.name].src);
      game.assets[asset.name].onloadend = (ev) => {
        queue.delete(ev.target.src);
      };
    }

    // TODO: wait for queue to finish
    let interval = setInterval(() => {
      if (queue.size === 0) {
        game.start();
        clearInterval(interval);
      }
    }, 250);
  });
</script>

<div class="overlay">
  <input
    class="hz-rb"
    type="number"
    min={0}
    max={25}
    value={rbHz}
    on:change={(ev) => game.updateHz(parseInt(ev.currentTarget.value))}
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
    padding: 8px;
  }

  canvas {
    max-width: 99%;
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
