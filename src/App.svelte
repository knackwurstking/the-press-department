<script>
  import { onMount } from "svelte";

  import Game from "./js/game";

  /** @type {HTMLCanvasElement} */
  let canvas;

  /** @type {HTMLImageElement} */
  let assetRolleLeft;
  /** @type {HTMLImageElement} */
  let assetRolleRight;

  /** @type {number} */
  let hzRB = 12;

  onMount(() => {
    const ctx = canvas.getContext("2d");
    canvas.width = 1730;
    canvas.height = 300;

    const game = new Game(canvas, canvas.width, canvas.height, hzRB, {
      rolleLeft: assetRolleLeft,
      rolleRight: assetRolleRight,
    });

    //let lastFrame = 0 - 600 / 12;
    (function animate(frame) {
      //if (frame - lastFrame >= 600 / 12) {
      //  game.draw(ctx, (lastFrame = frame));
      //}
      game.updateHz(hzRB);
      game.draw(ctx, frame);

      requestAnimationFrame(animate);
    })(0);
  });
</script>

<main>
  <div class="overlay">
    <input
      class="hz-rb"
      type="number"
      min={0}
      max={25}
      value={hzRB}
      on:change={(ev) => (hzRB = parseInt(ev.currentTarget.value))}
    />
  </div>
  <canvas bind:this={canvas} />
</main>

<img
  bind:this={assetRolleLeft}
  id="rolleLeft"
  src="assets/rolle-left_v2_6x296.png"
  alt="rolle"
/>
<img
  bind:this={assetRolleRight}
  id="rolleRight"
  src="assets/rolle-right_v2_6x296.png"
  alt="rolle"
/>
<img
  id="rbGestellAluBlock"
  src="assets/rb-gestell-alu-block_10x10.png"
  alt="gestell alu block"
/>
<img
  id="rbGestellQuer"
  src="assets/rb-gestell-quer_10x300.png"
  alt="gestell quer"
/>

<style>
  img {
    display: none;
  }

  main {
    width: 100vw;
    height: 100vh;
    display: flex;
    place-items: center;
    justify-content: center;
    background-color: rgb(39, 39, 39);
    overflow: auto;
  }

  canvas {
    /*
    border: var(--border);
    width: 1730px;
    height: 300px;
    */
    max-width: 99%;
  }

  .overlay {
    position: absolute;
    top: 0;
    right: 0;
    width: 5em;
    height: 3em;
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
