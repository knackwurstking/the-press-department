import Data from "./data";
import { EngineRollenBahn } from "./engine";

/**
 * @typedef Assets
 * @type {{
 *  rolleLeft?: HTMLImageElement,
 *  rolleRight?: HTMLImageElement,
 *  rbAluBlockLeft?: HTMLImageElement,
 *  rbAluBlockRight?: HTMLImageElement,
 *  rbRiemen290x5?: HTMLImageElement,
 *  rbRiemen270x5?: HTMLImageElement,
 * }}
 */

export default class Game {
  /**
   * @param {HTMLCanvasElement} canvas
   * @param {CanvasRenderingContext2D} ctx
   * @param {number} width
   * @param {number} height
   * @param {number} hz
   */
  constructor(canvas, ctx, width, height, hz) {
    this.canvas = canvas;
    this.ctx = ctx;
    this.width = width;
    this.height = height;

    this.updateHz(hz);
    this._lastFrame = 0 - this._fps;
    this._engineFrame = -1;

    /** @type {Assets} */
    this.assets = {};

    /** @type {EngineRollenBahn[]} */
    this.engines;

    console.log("[App.svelte] canvas size (x, y):", width, height);
  }

  initialize() {
    this.engines = []; // left to right

    let lastX = 0;
    for (let section of Data.rb) {
      let sX = lastX;
      let sY = 0;
      let width = section.engine.count * this.assets.rbAluBlockLeft.width;
      let height = this.canvas.height;
      lastX += width;

      this.engines.push(
        new EngineRollenBahn(
          section.name,
          section.engine.side,
          section.engine.count,
          sX,
          sY,
          width,
          height,
          this.assets
        )
      );
    }
  }

  /** @param {number} hz */
  updateHz(hz) {
    this._fps = 600 / hz;
  }

  /**
   * @param {number} frame
   */
  draw(frame) {
    if (frame - this._lastFrame >= this._fps) {
      this._engineFrame += 1;
      this._lastFrame = frame;

      if (this._engineFrame > 5) {
        this._engineFrame = 0;
      }

      for (let engine of this.engines) {
        engine.draw(this.ctx, this._engineFrame);
      }
    }
  }

  async start() {
    this.initialize();
    const animate = (/** @type {number} */ frame) => {
      this.draw(frame);
      requestAnimationFrame(animate);
    };

    animate(0);
  }
}
