import Data from "./data";
import { EngineRollenBahn } from "./Engine";

export default class Game {
  /**
   * @param {HTMLCanvasElement} canvas
   * @param {CanvasRenderingContext2D} ctx
   * @param {number} width
   * @param {number} height
   * @param {number} hz
   * @param {{
   *  rolleLeft: HTMLImageElement,
   *  rolleRight: HTMLImageElement,
   * }} assets
   */
  constructor(canvas, ctx, width, height, hz, assets) {
    this.canvas = canvas;
    this.ctx = ctx;
    this.width = width;
    this.height = height;
    this.assets = assets;

    this.updateHz(hz);
    this._lastFrame = 0 - this._fps;
    this._engineFrame = 0;

    // initialize
    this.initialize();

    console.log("[App.svelte] canvas size (x, y):", width, height);
  }

  initialize() {
    /** @type {EngineRollenBahn[]} */
    this.engines = []; // left to right
    let lastX = 0;
    for (let section of Data.rb) {
      let sX = lastX;
      let sY = 2;
      let width = section.engine.count * 10;
      let height = this.canvas.height;
      lastX += width;

      this.engines.push(
        new EngineRollenBahn(
          this.assets,
          section.engine.side,
          section.engine.count,
          sX,
          sY,
          width,
          height
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

      if (this._engineFrame > 6) {
        this._engineFrame = 1;
      }
    }

    for (let engine of this.engines) {
      engine.draw(this.ctx, this._engineFrame);
    }
  }
}
