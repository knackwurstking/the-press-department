<script>
  import { onMount } from "svelte";

  import Game from "./js/game";

  /** @type {HTMLCanvasElement} */
  let canvas;

  /** @type {HTMLImageElement} */
  let rolleLeft;
  /** @type {HTMLImageElement} */
  let rolleRight;
  /** @type {HTMLImageElement} */
  let rbAluBlockLeft;
  /** @type {HTMLImageElement} */
  let rbAluBlockRight;
  /** @type {HTMLImageElement} */
  let rbRiemen150x5;
  /** @type {HTMLImageElement} */
  let rbRiemen140x5;
  /** @type {HTMLImageElement} */
  let rbRiemen220x5;
  /** @type {HTMLImageElement} */
  let rbRiemen260x5;
  /** @type {HTMLImageElement} */
  let rbRiemen160x5;

  // some initial stuff
  let rbHz = 12;

  /** @type {Game} */
  let game;

  onMount(() => {
    const ctx = canvas.getContext("2d");
    canvas.width = 3460;
    canvas.height = 312;

    game = new Game(canvas, ctx, canvas.width, canvas.height, rbHz, {
      rolleLeft,
      rolleRight,
      rbAluBlockLeft,
      rbAluBlockRight,
      rbRiemen150x5,
      rbRiemen140x5,
      rbRiemen220x5,
      rbRiemen260x5,
      rbRiemen160x5,
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

<img
  bind:this={rolleLeft}
  src="assets/RolleLeft_6x296.png"
  alt="rolle"
/>
<img
  bind:this={rolleRight}
  src="assets/RolleRight_6x296.png"
  alt="rolle"
/>
<img
  bind:this={rbAluBlockLeft}
  src="assets/RollenBahnAluBlockLeft_20x10.png"
  alt="gestell alu block"
/>
<img
  bind:this={rbAluBlockRight}
  src="assets/RollenBahnAluBlockRight_20x10.png"
  alt="gestell alu block"
/>
<img
  bind:this={rbRiemen150x5}
  src="assets/RollenBahnRiemen-750_150x5-Sheetv2.png"
  alt="riemen"
/>
<img
  bind:this={rbRiemen140x5}
  src="assets/RollenBahnRiemen-700_140x5-Sheetv2.png"
  alt="riemen"
/>
<img
  bind:this={rbRiemen220x5}
  src="assets/RollenBahnRiemen-1100_220x5-Sheetv2.png"
  alt="riemen"
/>
<img
  bind:this={rbRiemen260x5}
  src="assets/RollenBahnRiemen-1300_260x5-Sheetv2.png"
  alt="riemen"
/>
<img
  bind:this={rbRiemen160x5}
  src="assets/RollenBahnRiemen-800_160x5-Sheetv2.png"
  alt="riemen"
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
    overflow: auto;
    background-image: url("./Ground_248x248.png"), url("/assets/Ground_248x248.png");
    padding: 8px;
  }

  canvas {
    width: 3460px;
    height: 312px;
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
