import Data from "./data";

export default class Game {
  /**
   * @param {number} width
   * @param {number} height
   * @param {number} hz
   * @param {{
   * rolleLeft: HTMLImageElement,
   * rolleRight: HTMLImageElement,
   * }} assets
   */
  constructor(width, height, hz, assets) {
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

          let image;
          if (engine.side === "right") {
            image = this.assets.rolleRight;
          } else if (engine.side === "left") {
            image = this.assets.rolleLeft;
          } else {
            continue;
          }

          ctx.drawImage(
            image,
            6 * this._engineFrame,
            0,
            6,
            image.height,
            posX,
            2,
            6,
            image.height
          );
        }
      }
    }
  }
}
