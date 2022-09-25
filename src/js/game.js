import Data from "./data";
import { EngineRollenBahn } from "./Engine";

export default class Game {
  /**
   * @param {HTMLCanvasElement} canvas
   * @param {number} width
   * @param {number} height
   * @param {number} hz
   * @param {{
   *  rolleLeft: HTMLImageElement,
   *  rolleRight: HTMLImageElement,
   * }} assets
   */
  constructor(canvas, width, height, hz, assets) {
    this.canvas = canvas;
    this.width = width;
    this.height = height;
    this.assets = assets;

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

    let lastX = 0;
    for (let section of Data.rb) {
      let sX = lastX;
      let sY = 2;
      let width = section.engine.count * 10;
      let height = this.canvas.height;
      lastX += width;

      // TODO: initialy cereate all engines and just draw here
      const erb = new EngineRollenBahn(
        this.assets,
        section.engine.side,
        section.engine.count,
        sX,
        sY,
        width,
        height
      );

      erb.draw(ctx, this._engineFrame);
    }
  }
}
