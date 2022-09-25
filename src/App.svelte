<script>
  import { onMount } from "svelte";

  import Game from "./js/game";

  /** @type {HTMLCanvasElement} */
  let canvas;

  /** @type {HTMLImageElement} */
  let assetRolleLeft;
  /** @type {HTMLImageElement} */
  let assetRolleRight;
  /** @type {HTMLImageElement} */
  let assetRBGestellAluBlockLeft;
  /** @type {HTMLImageElement} */
  let assetRBGestellAluBlockRight;

  // some initial stuff
  let rbHz = 12;

  /** @type {Game} */
  let game;

  onMount(() => {
    const ctx = canvas.getContext("2d");
    canvas.width = 1730;
    canvas.height = 312;

    game = new Game(canvas, ctx, canvas.width, canvas.height, rbHz, {
      rolleLeft: assetRolleLeft,
      rolleRight: assetRolleRight,
      rbGestellAluBlockLeft: assetRBGestellAluBlockLeft,
      rbGestellAluBlockRight: assetRBGestellAluBlockRight,
    });

    //let lastFrame = 0 - 600 / 12;
    (function animate(frame) {
      //if (frame - lastFrame >= 600 / 12) {
      //  game.draw(ctx, (lastFrame = frame));
      //}
      game.draw(frame);
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
      value={rbHz}
      on:change={(ev) => game.updateHz(parseInt(ev.currentTarget.value))}
    />
  </div>
  <canvas bind:this={canvas} />
</main>

<img
  bind:this={assetRolleLeft}
  id="rolleLeft"
  src="assets/RolleLeft_6x296.png"
  alt="rolle"
/>
<img
  bind:this={assetRolleRight}
  id="rolleRight"
  src="assets/RolleRight_6x296.png"
  alt="rolle"
/>
<img
  bind:this={assetRBGestellAluBlockLeft}
  id="rbGestellAluBlock"
  src="assets/RollenBahnAluBlockLeft_10x10.png"
  alt="gestell alu block"
/>
<img
  bind:this={assetRBGestellAluBlockRight}
  id="rbGestellAluBlock"
  src="assets/RollenBahnAluBlockRight_10x10.png"
  alt="gestell alu block"
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
    overflow: auto;
    background-image: url("./Ground_248x248.png"), url("/assets/Ground_248x248.png");
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
