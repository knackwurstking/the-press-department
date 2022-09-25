<script>
  import { onMount } from "svelte";

  import Data from "./js/data";

  /** @type {HTMLCanvasElement} */
  let canvas;

  /** @type {HTMLImageElement} */
  let assetRolleLeft;

  /** @type {number} */
  let hzRB = 12;

  class Game {
    /**
     * @param {number} width
     * @param {number} height
     * @param {number} hz
     */
    constructor(width, height, hz) {
      this.width = width;
      this.height = height;

      this.updateHz(hz);
      this._lastFrame = 0 - this._fps;
      this._engineFrame = 0;

      console.log("[App.svelte] canvas size (x, y):", width, height);
    }

    /** @param {number} hz */
    updateHz(hz) {
      this._fps = 600 / hz;
    }

    /**
     * @param {CanvasRenderingContext2D} ctx
     * @param {number} frame
     */
    draw(ctx, frame) {
      if (frame - this._lastFrame >= this._fps) {
        this._engineFrame += 1;
        this._lastFrame = frame;

        if (this._engineFrame > 6) {
          this._engineFrame = 1;
        }
      }

      let index = -1;
      for (let section of Data.rb) {
        const engine = section.engine;

        if (engine.type === Data.ROLLEN_GRIP) {
          const assetRolleWidth = 10;

          for (let x = 0; x < engine.count; x++) {
            index += 1;

            let posX = index * assetRolleWidth + 2;
            if (posX >= this.width) {
              throw `posX for "rolle" is out of range (${posX})`;
            }

            ctx.drawImage(
              assetRolleLeft,
              6 * this._engineFrame,
              0,
              6,
              assetRolleLeft.height,
              posX,
              2,
              6,
              assetRolleLeft.height
            );
          }
        }
      }
    }
  }

  onMount(() => {
    const ctx = canvas.getContext("2d");
    canvas.width = 1730;
    canvas.height = 300;

    const game = new Game(canvas.width, canvas.height, hzRB);

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
  src="assets/rolle-left_6x296.png"
  alt="iamge"
/>
<img id="rolleRight" src="assets/rolle-right_6x296.png" alt="iamge" />
<img
  id="rbGestellAluBlock"
  src="assets/rb-gestell-alu-block_10x10.png"
  alt="iamge"
/>
<img id="rbGestellQuer" src="assets/rb-gestell-quer_10x300.png" alt="iamge" />

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
  }

  canvas {
    border: var(--border);
    width: 1730px;
    height: 300px;
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
